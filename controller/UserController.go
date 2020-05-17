package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-blog/cache"
	"go-blog/constant"
	"go-blog/dao"
	"go-blog/dto"
	"go-blog/model"
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

	var userInfo = model.UserInfo{}

	err := ctx.BindJSON(&userInfo)

	var result interface{}
	if err != nil {
		log.Println("解析登陆信息失败")
		result = dto.BaseResult{
			Code: constant.FAIL.Code,
			Msg:  "解析登陆信息失败",
		}
	}
	var loginId = userInfo.LoginId.String
	var pass = userInfo.Password.String

	var userInfoQuery *model.UserInfo
	userInfoQuery, err = dao.PublicUserInfoDao.SelectByLoginId(loginId)

	if err != nil {

		log.Println("查询数据出错")
		result = dto.BaseResult{
			Code: constant.FAIL.Code,
			Msg:  "查询数据出错,请稍后重试",
		}
		ctx.JSON(http.StatusOK, result)
		return
	}

	if userInfoQuery != nil {

		if userInfoQuery.Password.String == pass {

			token := uuid.NewV4().String()
			var userInfoBytes []byte
			userInfoBytes,err = json.Marshal(userInfoQuery)
			if err != nil{

				result = dto.BaseResult{
					Code: constant.FAIL.Code,
					Msg:  constant.FAIL.Msg,
				}
				ctx.JSON(http.StatusOK, result)
				return
			}
			cache.CacheUtil.Set([]byte(token), userInfoBytes, constant.SESSION_EXPORE_SECONDS)
			result = dto.BaseResult{
				Code: constant.SUCCESS.Code,
				Msg:  constant.SUCCESS.Msg,
				Data: token,
			}

			ctx.JSON(http.StatusOK, result)
			return
		}
	}

	ctx.JSON(http.StatusOK, dto.BaseResult{
		Code: constant.LOGIN_FAIL.Code,
		Msg:  constant.LOGIN_FAIL.Msg,
	})
}

func GetLoginInfo(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")

	tokenBytes,err := cache.CacheUtil.Get([]byte(token))

	if err != nil{
		ctx.JSON(http.StatusOK,dto.BaseResult{
			Code: constant.FORBIDDEN.Code,
			Msg:  constant.FORBIDDEN.Msg,
		})

		return
	}
	loginInfo := model.UserInfo{}
	err = json.Unmarshal(tokenBytes,&loginInfo)
	if err != nil{
		ctx.JSON(http.StatusOK,dto.BaseResult{
			Code: constant.FAIL.Code,
			Msg:  constant.FAIL.Msg,
		})

		return
	}

	ctx.JSON(http.StatusOK,dto.BaseResult{
		Code: constant.SUCCESS.Code,
		Msg:  constant.SUCCESS.Msg,
		Data: &loginInfo,
	})
}
