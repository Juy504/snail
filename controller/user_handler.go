package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	repo "snail/repository"
	"snail/serializer"
	"snail/service"
	"snail/util"
)

func AuthHandler(c *gin.Context) {
	service.Auth(c)
}

func JWTAuthHandler(c *gin.Context) {
	service.JWTAuth(c)
}

func UserInfoHandler(c *gin.Context) {
	userid := c.Query("user_id")
	userType:= new(service.UserService)
	var user repo.UserService
	user = userType
	userinfo, err := user.GetUserInfo(c, userid, "")
	if err != nil{
		err := serializer.ErrorData(err, 1)
		c.JSON(http.StatusNotFound, err)
		return
	}
	res := serializer.SuccessData(userinfo)
	c.JSON(http.StatusOK, res)
	return
}

func PhoneLogin(c *gin.Context){
	phone := c.Query("phone")
	password := c.Query("password")
	userType:= new(service.UserService)
	var user repo.UserService
	user = userType
	userinfo, err := user.GetUserInfo(c, "", phone)
	if err != nil{
		err := serializer.ErrorData(err, 1)
		c.JSON(http.StatusNotFound, err)
		return
	}
	hashcode, salt := userinfo.Password, userinfo.Salt
	//校验用户密码
	if ok := util.VerifyPassword(password, hashcode, salt); !ok{
		err := serializer.ErrorData(errors.New("密码错误"), 1)
		c.JSON(http.StatusOK, err)
		return
	}
	res := serializer.SuccessData(userinfo)
	c.JSON(http.StatusOK, res)
	return
}

func Regist(c *gin.Context){
	phone := c.Query("phone")
	password := c.Query("password")
	if phone == "" || password == ""{
		err := serializer.ErrorData(errors.New("参数不合法"), 1)
		c.JSON(http.StatusOK, err)
		return
	}
	//userType:= new(service.UserService)
	//var user repo.UserService
	//user = userType
}