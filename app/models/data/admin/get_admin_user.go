package admin

import (
	"github.com/layasugar/laya-template/handle/model/dao"
	"github.com/layasugar/laya-template/handle/model/dao/db"
)

type (
	AdminUser = db.AdminUser
)

// GetUserInfoByUsername 根据用户名获取用户信息
func GetUserInfoByUsername(ctx *core.Context, username string) (*AdminUser, error) {
	var u AdminUser
	err := dao.Orm(ctx).WithContext(ctx).Where("username = ?", username).Find(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}
