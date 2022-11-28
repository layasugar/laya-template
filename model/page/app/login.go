package app

import (
	"errors"
	"fmt"

	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/model/data/user"
)

type LoginParam struct {
	Mobile  string `json:"mobile" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Scene   string `json:"scene" binding:"required,oneof=both login register"` // 场景值 这里是 login/register
	Channel string `json:"channel" binding:"omitempty"`                        // 注册渠道
}

type LoginResponse struct {
	UserId uint64 `json:"user_id"`
	Mobile string `json:"mobile"`
	Token  string `json:"token"`
}

// Login 用户登录
func Login(ctx *laya.Context, pm *LoginParam) (*LoginResponse, error) {
	// 验证验证码
	err := checkSmsCode(ctx, pm)
	if err != nil {
		ctx.Warn("Login checkSmsCode fail,err:%s", err.Error())
		return nil, err
	}

	// 获取用户
	userInfo, err := userOperation(ctx, pm)
	if err != nil {
		return nil, err
	}

	// 登录
	token, err := user.GetToken(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	resp := LoginResponse{
		UserId: userInfo.ID,
		Mobile: userInfo.Mobile,
		Token:  token,
	}

	return &resp, err
}

// userOperation 用户操作
func userOperation(ctx *laya.Context, pm *LoginParam) (*user.User, error) {
	userinfo, err := user.GetUserByMobile(ctx, pm.Mobile)
	if err != nil {
		return nil, err
	}

	if userinfo == nil || userinfo.ID <= 0 {
		userReq := user.User{
			Username: pm.Mobile,
			Mobile:   pm.Mobile,
			Nickname: fmt.Sprintf("用户%s", pm.Mobile[7:]),
			Avatar:   "",
		}
		errCreateUser := user.CreateUser(ctx, &userReq)
		if errCreateUser != nil {
			return nil, errCreateUser
		}

		userinfo = &userReq
	}

	return userinfo, nil
}

//checkSmsCode 检查验证码
func checkSmsCode(ctx *laya.Context, pms *LoginParam) error {
	smsCode, err := user.GetSmsCode(ctx, pms.Code, pms.Mobile)
	if err != nil {
		ctx.Warn("checkSmsCode getSmsCode fail,err:%s", err.Error())
		return err
	}

	// 判断验证码
	if smsCode != pms.Code {
		return errors.New("短信验证码错误")
	}

	// 删除redis smsCode
	err = user.DelSmsCode(ctx, pms.Code, pms.Mobile)
	if err != nil {
		ctx.Warn("checkSmsCode delSmsCode fail,err:%s", err.Error())
		return err
	}

	return nil
}

// Logout 退出登录
func Logout(ctx *laya.Context) error {
	token := ctx.Gin().GetHeader(global.UserAuth)
	return user.DelToken(ctx, token)
}
