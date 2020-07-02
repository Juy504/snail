package controller

import (
	"github.com/gin-gonic/gin"
	"snail/service"
)

func AuthHandler(c *gin.Context) {
	service.Auth(c)
}
func JWTAuthHandler(c *gin.Context) {
	service.JWTAuth(c)
}

func UserInfoHandler(c *gin.Context) {
	service.UserInfo(c)
}
