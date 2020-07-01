package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Models "snail/model"
	"snail/util"
	"strings"
)

func Auth(c *gin.Context) {
	name, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	// 用户发送用户名和密码过来
	var user Models.User
	//查询数据库
	user.ID = 1
	user.Username = "admin"
	user.Password = "admin"
	// 校验用户名和密码是否正确
	if user.Username == name && user.Password == password {
		// 生成Token
		tokenString, _ := util.GenToken(&user)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定??
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("id", mc.Id)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}

}

/**
获取当前用户信息
*/
func UserInfo(c *gin.Context) {
	username := c.MustGet("username").(string)
	id := c.MustGet("id").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"id": id, "username": username},
	})
}
