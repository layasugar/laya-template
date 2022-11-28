package user

import (
	"encoding/json"
	"fmt"

	"github.com/layasugar/laya"
	"github.com/layasugar/laya/gcnf"

	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/model/dao"
	"github.com/layasugar/laya-template/utils"
)

// GetToken 获取token
func GetToken(ctx *laya.Context, userInfo *User) (string, error) {
	token := utils.RandToken()
	tokenRedisKey := gcnf.AppName() + global.TokenRedisKey
	key := fmt.Sprintf(tokenRedisKey, token)
	userData, err := json.Marshal(userInfo)
	if err != nil {
		ctx.Warn("GetToken json.Marshal user fail, err: %s", err.Error())
		return "", err
	}
	err = dao.Rdb().SetEX(ctx, key, userData, global.TokenExpire).Err()
	if err != nil {
		ctx.Warn("GetToken SetEx fail, err: %s", err.Error())
		return "", err
	}

	return token, nil
}

// DelToken redis清理token
func DelToken(ctx *laya.Context, token string) error {
	tokenRedisKey := gcnf.AppName() + global.TokenRedisKey
	key := fmt.Sprintf(tokenRedisKey, token)
	err := dao.Rdb().Del(ctx, key).Err()
	if err != nil {
		ctx.Warn("Logout rdb.del fail, err: %s", err.Error())
		return err
	}
	return nil
}
