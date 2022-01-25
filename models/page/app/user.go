package app

import (
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/models/data"
	"gorm.io/gorm"
)

type UserParam struct {
	Id uint64 `json:"id" binding:"required"`
}

func GetUserInfo(ctx *laya.WebContext, param *UserParam) (interface{}, error) {
	user, err := data.GetUserById(ctx, param.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Info("用户中心", "用户不存在")
		return nil, global.UserNotFound
	}
	return user, nil
}
