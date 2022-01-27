package user

import (
	"encoding/json"
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/utils"
	"github.com/layasugar/laya/genv"
)

// GetToken 获取token
func GetToken(ctx *laya.WebContext, userInfo *User) (string, error) {
	token := utils.RandToken()
	tokenRedisKey := genv.AppName() + global.TokenRedisKey
	key := fmt.Sprintf(tokenRedisKey, token)
	userData, err := json.Marshal(userInfo)
	if err != nil {
		ctx.WarnF("GetToken json.Marshal user fail, err: %s", err.Error())
		return "", err
	}
	err = dao.Rdb.SetEX(ctx, key, userData, global.TokenExpire).Err()
	if err != nil {
		ctx.WarnF("GetToken SetEx fail, err: %s", err.Error())
		return "", err
	}

	return token, nil
}

// DelToken redis清理token
func DelToken(ctx *laya.WebContext, token string) error {
	tokenRedisKey := genv.AppName() + global.TokenRedisKey
	key := fmt.Sprintf(tokenRedisKey, token)
	err := dao.Rdb.Del(ctx, key).Err()
	if err != nil {
		ctx.WarnF("Logout rdb.del fail, err: %s", err.Error())
		return err
	}
	return nil
}
