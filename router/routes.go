package router

import (
	"github.com/gin-gonic/gin"
	"snail/controller"
	"snail/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Logger())

	//jwt鉴权
	r.POST("auth", controller.AuthHandler)

	//校验
	r.Use(controller.JWTAuthHandler)
	//获取当前用户信息
	user := r.Group("user")
	{
		user.GET("userInfo", controller.UserInfoHandler)
	}
	return r
}
