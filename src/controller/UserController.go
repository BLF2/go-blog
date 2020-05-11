package controller

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-blog/src/cache"
	"go-blog/src/constant"
	"go-blog/src/dao"
	"go-blog/src/dto"
	"go-blog/src/model"
	"gopkg.in/guregu/null.v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Insert(context *gin.Context) {

	var userInfo = model.UserInfo{}
	context.BindJSON(&userInfo)

	userInfo.CreateTime = time.Now()
	userInfo.CreateUser = null.StringFrom("admin")

	id, err := dao.PublicUserInfoDao.Insert(&userInfo)

	var result dto.BaseResult

	if err != nil {
		log.Println("插入数据出现错误", err)
		result = dto.BaseResult{
			Code: constant.FAIL.Code,
			Msg:  constant.FAIL.Msg,
		}
	} else {
		log.Println("插入数据成功")
		result = dto.BaseResult{
			Code: constant.SUCCESS.Code,
			Msg:  constant.SUCCESS.Msg,
			Data: id,
		}
	}

	context.JSON(http.StatusOK, result)
}

func GetById(context *gin.Context) {

	var id int64
	var err error
	var result dto.BaseResult
	id, err = strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		log.Println("根据id查询用户信息失败  id不可转换为int64")
		result = dto.BaseResult{
			Code: constant.FAIL.Code,
			Msg:  constant.FAIL.Msg,
			Data: "id不可转换为int64",
		}
	} else {
		var userInfo *model.UserInfo
		userInfo, err = dao.PublicUserInfoDao.SelectById(id)
		if err != nil {
			log.Println("根据id查询用户信息失败，数据库错误")
			result = dto.BaseResult{
				Code: constant.FAIL.Code,
				Msg:  constant.FAIL.Msg,
				Data: "数据库错误",
			}
		} else {
			result = dto.BaseResult{
				Code: constant.SUCCESS.Code,
				Msg:  constant.SUCCESS.Msg,
				Data: userInfo,
			}
		}

	}

	context.JSON(http.StatusOK, result)
}

func Login(ctx *gin.Context) {

	var userInfo  = model.UserInfo{}

	err := ctx.BindJSON(&userInfo)

	var result interface{}
	if err != nil{
		log.Println("解析登陆信息失败")
		result = dto.BaseResult{
			Code: constant.FAIL.Code,
			Msg: "解析登陆信息失败",
		}
	}else{
		var loginId = userInfo.LoginId.String
		var pass = userInfo.Password.String

		var userInfoQuery *model.UserInfo
		userInfoQuery,err= dao.PublicUserInfoDao.SelectByLoginId(loginId)

		if err != nil {
			
			log.Println("查询数据出错")
			result = dto.BaseResult{
				Code: constant.FAIL.Code,
				Msg:  "查询数据出错,请稍后重试",
			}
			ctx.JSON(http.StatusOK,result)
			return
		}

		if userInfoQuery != nil{

			if userInfoQuery.Password.String == pass{

				token := uuid.NewV4().String()
				cache.CacheUtil.Set([]byte(token),[]byte(loginId),constant.SESSION_EXPORE_SECONDS)
				result = dto.BaseResult{
					Code: constant.SUCCESS.Code,
					Msg:  constant.SUCCESS.Msg,
					Data: token,
				}

				ctx.JSON(http.StatusOK,result)
				return
			}
		}

		ctx.JSON(http.StatusOK,dto.BaseResult{
			Code: constant.SUCCESS.Code,
			Msg:  "用户名不存在或者用户名与密码不匹配",
		})
	}

	ctx.JSON(http.StatusOK,result)
}
