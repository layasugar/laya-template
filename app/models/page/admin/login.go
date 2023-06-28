package admin

import (
	"crypto/md5"
	"fmt"
	admin2 "github.com/layasugar/laya-template/handle/model/data/admin"
	"io"

	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/global/errno"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp = struct {
	Id     uint64 `json:"id"`
	Role   uint64 `json:"role"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Token  string `json:"token"`
}

func Login(ctx *core.Context, request *LoginReq) (*LoginRsp, error) {
	userinfo, err := admin2.GetUserInfoByUsername(ctx, request.Username)
	if err != nil {
		ctx.Info("Login get user info error: ", err)
		return nil, err
	}
	if userinfo == nil || userinfo.ID <= 0 {
		return nil, errno.UserNotFound
	}
	h := md5.New()
	_, _ = io.WriteString(h, request.Password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	if userinfo.Password != pwmd5 {
		ctx.Info("Login password is wrong", nil)
		return nil, errno.UserNotFound
	}

	token, err := admin2.GetToken(ctx, userinfo)
	if err != nil {
		return nil, err
	}

	return &LoginRsp{
		Id:     userinfo.ID,
		Role:   userinfo.DefaultRole,
		Name:   userinfo.Username,
		Avatar: userinfo.Avatar,
		Token:  token,
	}, err
}

// Logout 退出登录
func Logout(ctx *core.Context) error {
	token := ctx.Gin().GetHeader(global.UserAuth)
	return admin2.DelToken(ctx, token)
}
