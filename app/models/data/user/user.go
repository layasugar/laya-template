package user

import (
	"errors"
	"github.com/layasugar/laya-template/app/models/dao"
	"github.com/layasugar/laya-template/app/models/dao/db"
	"github.com/layasugar/laya-template/pkg/core"

	"gorm.io/gorm"
)

type (
	User              = db.User
	GetUserListParams struct {
		Page     uint
		PageSize uint
	}
)

func GetUserById(ctx *core.Context, userId uint64) (*User, error) {
	var u User
	var orm = dao.Orm(ctx)

	err := orm.Model(&User{}).Where("id = ?", userId).First(&u).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Warn("GetUserById, err: %s", err.Error())
		return nil, err
	}
	return &u, nil
}

func GetUserListByZone(ctx *core.Context, Id int64) ([]User, error) {
	var users []User
	err := dao.Orm(ctx).Model(&User{}).Where("id = ?", Id).Find(&users).Error
	if err != nil {
		ctx.Warn("GetUserListByZone, err: %s", err.Error())
		return nil, err
	}
	return users, err
}

func GetUserByMobile(ctx *core.Context, mobile string) (*User, error) {
	var u User
	err := dao.Orm(ctx).WithContext(ctx).Where("mobile = ?", mobile).First(&u).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Warn("GetUserByMobile, err: %s", err.Error())
		return nil, err
	}
	return &u, nil
}

func CreateUser(ctx *core.Context, u *User) error {
	err := dao.Orm(ctx).WithContext(ctx).Create(u).Error
	if err != nil {
		ctx.Warn("CreateUser fail, err:%s", err.Error())
		return err
	}
	return nil
}
