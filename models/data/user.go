package data

import (
	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya-go/models/dao"
	"github.com/layasugar/laya-go/models/dao/db"
)

type User = db.User

func GetUserById(c *gin.Context, userId uint64) (*User, error) {
	var u User
	err := dao.DB.Model(&User{}).Where("id = ?", userId).First(&u).Error
	return &u, err
}

func GetUserListByZone(c *gin.Context, Id int64) ([]*User, error) {
	var users []*User
	err := dao.DB.Model(&User{}).Where("id = ?", Id).Find(&users).Error
	return users, err
}
