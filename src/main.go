package main

import (
	"github.com/gin-gonic/gin"
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

	router.Run("0.0.0.0:8080")
}

