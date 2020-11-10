package repository

import (
	"github.com/gin-gonic/gin"
	"snail/model"
)

type UserService interface {
	GetUserInfo(c *gin.Context, uid string, phone string) (*model.User, error)
	GetByUserIds(c *gin.Context, uids string) (*[]model.User, error)
	Create(c *gin.Context, u *model.User) (int64, error)
	Update(c *gin.Context, u *model.User) (*model.User, error)
}
