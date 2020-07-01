package controller

import (
	"github.com/gin-gonic/gin"
	"snail/service"
)

func AuthHandler(c *gin.Context) {
	service.Auth(c)
}
func JWTAuthHandler() {
	service.JWTAuth()
}

func UserInfoHandler(c *gin.Context) {
	service.UserInfo(c)
}
