package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	repo "snail/repository"
	"snail/serializer"
	"snail/service"
)

func AuthHandler(c *gin.Context) {
	service.Auth(c)
}

func JWTAuthHandler(c *gin.Context) {
	service.JWTAuth(c)
}

func UserInfoHandler(c *gin.Context) {
	userid := c.Query("user_id")
	fmt.Println(userid)
	userType:= new(service.UserService)
	var user repo.UserService
	user = userType
	userinfo, err := user.GetByUserId(c, userid)
	if err != nil{
		err := serializer.ErrorData(err, 1)
		c.JSON(http.StatusNotFound, err)
		return
	}
	res := serializer.SuccessData(userinfo)
	c.JSON(http.StatusOK, res)
	return
}
