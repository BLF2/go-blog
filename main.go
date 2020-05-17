package main

import (
	"github.com/gin-gonic/gin"
	_ "go-blog/cache"
	"go-blog/compont"
	"go-blog/controller"
	"go-blog/dto"
)

func main()  {

	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, dto.BaseResult{
			Code: "000000",
			Msg: "success",
			Data: "test",
		})
	})

	router.Static("/static", "./static")
	router.Use(compont.Auth)
	router.POST("/user/insert", controller.Insert)
	router.GET("/user/findById/:id", controller.GetById)
	router.POST("/user/login", controller.Login)
	router.GET("/user/getLoginInfo",controller.GetLoginInfo)

	router.Run(":8080")
}

