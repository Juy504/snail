package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	mw "snail/middleware"
	"net/http"
	Model "snail/model"
	repo "snail/repository"
	"snail/util"
	"strings"
)

func Auth(c *gin.Context) {
	name, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	// 用户发送用户名和密码过来
	var user Model.User
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
func JWTAuth(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
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

/**
获取当前用户信息
*/
type UserService struct{
	Repo repo.UserService `inject:""`
}
func (u *UserService)GetUserInfo(c *gin.Context, uid string, phone string) (*Model.User, error) {
	var user Model.User
	db := util.DB
	if uid == "" && phone == "" {
		mw.LogClient.WithFields(log.Fields{
			"user_id": uid,
			"phone": phone,
		}).Info("请求参数不合法", "info")
		return &Model.User{}, errors.New("参数不合法")
	}
	if uid != ""{
		db.Where("user_id = ?", uid).First(&user)
	}else if phone != ""{
		db.Where("phone = ?", phone).First(&user)
	}
	if user.ID == 0 {
		mw.LogClient.WithFields(log.Fields{
			"user_id": uid,
		}).Info("未找到用户信息", "info")
		return &Model.User{}, errors.New("未查到用户信息")
	}
	panic("LogClient")
	mw.LogClient.WithFields(log.Fields{
		"user_id": user.UserId,
	}).Info("获取用户信息成功", "info")
	return &Model.User{ID: user.ID, Nickname:user.Nickname, Username:user.Username, UserId:user.UserId,
		Title:user.Title, Salt: user.Salt}, nil
}

func (u *UserService)GetByUserIds(c *gin.Context, uids string) (*[]Model.User, error) {
	return u.Repo.GetByUserIds(c, uids)
}
func (u *UserService)Create(c *gin.Context, m *Model.User) (int64, error) {
	return u.Repo.Create(c, m)
}
func (u *UserService)Update(c *gin.Context, m *Model.User) (*Model.User, error) {
	return u.Repo.Update(c, m)
}
