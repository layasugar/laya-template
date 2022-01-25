package data

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/models/dao/db"
)

type User = db.User

func GetUserById(ctx *laya.WebContext, userId uint64) (*User, error) {
	var u User
	err := dao.DB.Model(&User{}).Where("id = ?", userId).First(&u).Error
	return &u, err
}

func GetUserListByZone(ctx *laya.WebContext, Id int64) ([]*User, error) {
	var users []*User
	err := dao.DB.Model(&User{}).Where("id = ?", Id).Find(&users).Error
	return users, err
}
