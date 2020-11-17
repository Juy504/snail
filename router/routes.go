package router

import (
	"github.com/gin-gonic/gin"
	handler "snail/controller"
	"snail/middleware"
	_ "snail/util"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Logger())

	//jwt鉴权
	//r.POST("auth", controller.AuthHandler)

	//校验
	//r.Use(controller.JWTAuthHandler)
	//获取当前用户信息
	api := r.Group("v1")
	{
		api.GET("userInfo", handler.UserInfoHandler)
		api.GET("phoneLogin", handler.PhoneLogin)
		api.GET("regist", handler.Regist)
	}
	return r
}
