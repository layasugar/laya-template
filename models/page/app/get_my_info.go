package app

import (
	"encoding/json"
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/global/errno"
	"github.com/layasugar/laya-template/models/data/user"
	"gorm.io/gorm"
)

type (
	GetMyInfoRsp struct {
		ID       uint64 `ddb:"id" json:"id"`
		Username string `ddb:"username" json:"username"` // 用户名
		Nickname string `ddb:"nickname" json:"nickname"` // 昵称
		Avatar   string `ddb:"avatar" json:"avatar"`     // 头像
		Mobile   string `ddb:"mobile" json:"mobile"`     // 手机号
		Status   uint8  `ddb:"status" json:"status"`     // 状态
		Channel  uint8  `ddb:"channel" json:"channel"`   // 注册渠道
	}
)

func GetUserInfo(ctx *laya.WebContext) (*GetMyInfoRsp, error) {
	if userInfoStr := ctx.GetString(global.UserInfo); userInfoStr != "" {
		var userInfo user.User
		err := json.Unmarshal([]byte(userInfoStr), &userInfo)
		if err != nil {
			return nil, err
		}

		u, err := user.GetUserById(ctx, userInfo.ID)
		if err != nil {
			return nil, err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.ErrorF("用户中心: %s", "用户不存在")
			return nil, errno.UserNotFound
		}
		res := GetMyInfoRsp{
			ID:       u.ID,
			Username: u.Username,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Mobile:   u.Mobile,
			Status:   u.Status,
			Channel:  u.Channel,
		}
		return &res, nil
	} else {
		return nil, errno.ComUnauthorized
	}
}
