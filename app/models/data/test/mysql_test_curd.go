package test

import (
	"github.com/layasugar/laya-template/app/models/dao"
	"github.com/layasugar/laya-template/app/models/dao/db"
	"github.com/layasugar/laya-template/pkg/core"
)

type (
	User = db.User
)

func MysqlUserCreate(ctx *core.Context, params User) error {
	return dao.Orm(ctx).Model(&User{}).Create(&params).Error
}

func MysqlUserUpdate(ctx *core.Context, phone string, nickName string) error {
	return dao.Orm(ctx).Model(&User{}).Where("mobile = ?", phone).Updates(map[string]interface{}{
		"nickname": nickName,
	}).Error
}

func MysqlUserSelect(ctx *core.Context, phone string) (*User, error) {
	var u = &User{}
	err := dao.Orm(ctx).Where("mobile = ?", phone).First(u).Error
	return u, err
}

func MysqlUserDel(ctx *core.Context, phone string) error {
	return dao.Orm(ctx).Where("mobile = ?", phone).Unscoped().Delete(&User{}).Error
}
