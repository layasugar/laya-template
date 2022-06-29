package admin

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/models/dao/db"
)

type (
	AdminUser = db.AdminUser
)

// GetUserInfoByUsername 根据用户名获取用户信息
func GetUserInfoByUsername(ctx *laya.WebContext, username string) (*AdminUser, error) {
	var u AdminUser
	err := dao.Orm(ctx).WithContext(ctx).Where("username = ?", username).Find(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}
