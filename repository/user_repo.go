package repository

import (
	"github.com/gin-gonic/gin"
	"snail/model"
)

type UserService interface {
	GetByUserId(c *gin.Context, uid string) (*model.User, error)
	GetByUserIds(c *gin.Context, uid string) (*[]model.User, error)
	Create(c *gin.Context, u *model.User) (int64, error)
	Update(c *gin.Context, u *model.User) (*model.User, error)
}
