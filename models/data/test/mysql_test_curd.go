package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/handle/model/dao"
	"github.com/layasugar/laya-template/handle/model/dao/db"
)

type (
	User = db.User
)

func MysqlUserCreate(ctx *laya.Context, params User) error {
	return dao.Orm(ctx).Model(&User{}).Create(&params).Error
}

func MysqlUserUpdate(ctx *laya.Context, phone string, nickName string) error {
	return dao.Orm(ctx).Model(&User{}).Where("mobile = ?", phone).Updates(map[string]interface{}{
		"nickname": nickName,
	}).Error
}

func MysqlUserSelect(ctx *laya.Context, phone string) (*User, error) {
	var u = &User{}
	err := dao.Orm(ctx).Where("mobile = ?", phone).First(u).Error
	return u, err
}

func MysqlUserDel(ctx *laya.Context, phone string) error {
	return dao.Orm(ctx).Where("mobile = ?", phone).Unscoped().Delete(&User{}).Error
}
