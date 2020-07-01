package router

import (
	"github.com/gin-gonic/gin"
	"snail/middleware"
	"snail/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Logger())

	//jwt鉴权
	r.POST("auth", service.Auth)

	//校验
	r.Use(service.JWTAuth())
	//获取当前用户信息
	user := r.Group("user")
	{
		user.GET("userInfo", service.UserInfo)
	}
	return r
}
