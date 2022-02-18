package user

import (
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/models/dao/db"
	"gorm.io/gorm"
)

type (
	User              = db.User
	GetUserListParams struct {
		Page     uint
		PageSize uint
	}
)

func GetUserById(ctx *laya.WebContext, userId uint64) (*User, error) {
	var u User
	var orm = dao.Orm(ctx)

	err := orm.Model(&User{}).Where("id = ?", userId).First(&u).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.WarnF("GetUserById, err: %s", err.Error())
		return nil, err
	}
	return &u, nil
}

func GetUserListByZone(ctx *laya.WebContext, Id int64) ([]User, error) {
	var users []User
	err := dao.DB.WithContext(ctx).Model(&User{}).Where("id = ?", Id).Find(&users).Error
	if err != nil {
		ctx.WarnF("GetUserListByZone, err: %s", err.Error())
		return nil, err
	}
	return users, err
}

func GetUserByMobile(ctx *laya.WebContext, mobile string) (*User, error) {
	var u User
	err := dao.DB.WithContext(ctx).Where("mobile = ?", mobile).First(&u).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.WarnF("GetUserByMobile, err: %s", err.Error())
		return nil, err
	}
	return &u, nil
}

func CreateUser(ctx *laya.WebContext, u *User) error {
	err := dao.DB.WithContext(ctx).Create(u).Error
	if err != nil {
		ctx.WarnF("CreateUser fail, err:%s", err.Error())
		return err
	}
	return nil
}
