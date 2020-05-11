package main

import (
	"github.com/gin-gonic/gin"
	_ "go-blog/src/cache"
	"go-blog/src/controller"
	"go-blog/src/dto"
)

func main()  {

	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200,dto.BaseResult{
			Code: "000000",
			Msg: "success",
			Data: "test",
		})
	})

	router.POST("/user/insert",controller.Insert)
	router.GET("/user/findById/:id",controller.GetById)
	router.POST("/user/login",controller.Login)

	router.Run("0.0.0.0:8080")
}

