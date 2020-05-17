package compont

import (
	"github.com/gin-gonic/gin"
	"go-blog/constant"
	"go-blog/dto"
	"log"
	"net/http"
	"strings"
)
import cache "go-blog/cache"

func Auth(ctx *gin.Context)  {

	log.Println("进入token验证")

	if strings.HasPrefix(ctx.Request.URL.Path,"/static/login.html") ||
		strings.HasPrefix(ctx.Request.URL.Path,"/user/login"){
		log.Println("登陆相关，不校验token")
		return
	}
	token := ctx.Request.Header.Get("token")

	if token == ""{
		result := dto.BaseResult{
			Code: constant.FORBIDDEN.Code,
			Msg:  constant.FORBIDDEN.Msg,
		}

		ctx.JSON(http.StatusFound,result)
		ctx.Abort()
		return
	}

	if _,err := cache.CacheUtil.Get([]byte(token)) ;err != nil{

		result := dto.BaseResult{
			Code: constant.FORBIDDEN.Code,
			Msg:  constant.FORBIDDEN.Msg,
			Data: "/static/home.html",
		}
		log.Println("token验证不通过")
		ctx.JSON(http.StatusFound,result)
		ctx.Abort()
		return
	}

	log.Println("token验证通过")
}
