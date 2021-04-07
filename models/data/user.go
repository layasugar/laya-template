package data

import (
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/models/dao"
	"github.com/layatips/laya-go/models/dao/db"
)

type User = db.User

func GetUserById(c *gin.Context, userId uint64) (*User, error) {
	var u User
	err := dao.DB.Model(&User{}).Where("id = ?", userId).First(&u).Error
	return &u, err
}

func GetUserListByZone(c *gin.Context, Zone string) ([]*User, error) {
	var users []*User
	err := dao.DB.Model(&User{}).Where("id = ?", Zone).Find(&users).Error
	return users, err
}
