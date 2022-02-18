package user

import (
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya/genv"
)

// GetSmsCode 获取短信验证码
func GetSmsCode(ctx *laya.WebContext, scene, mobile string) (string, error) {
	verifyCodeKey := genv.AppName() + global.VerifyCodeKey
	key := fmt.Sprintf(verifyCodeKey, scene, mobile)
	code, err := dao.Rdb().Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

// DelSmsCode 删除短信验证码
func DelSmsCode(ctx *laya.WebContext, scene, mobile string) error {
	verifyCodeKey := genv.AppName() + global.VerifyCodeKey
	key := fmt.Sprintf(verifyCodeKey, scene, mobile)
	err := dao.Rdb().Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
