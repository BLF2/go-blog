package controller

import (
	"github.com/gin-gonic/gin"
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
