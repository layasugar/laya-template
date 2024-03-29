package admin

import (
	"github.com/layasugar/laya-template/app/models/data/admin"
	"github.com/layasugar/laya-template/global/pagination"
	"github.com/layasugar/laya-template/pkg/core"
)

type (
	GetUserListReq struct {
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
		Id       uint64 `json:"id"`
		Username string `json:"username"`
		Status   uint8  `json:"status"`
		Mobile   string `json:"mobile"`
	}
	GetUserListRsp struct {
		List       []User                `json:"list"`
		Pagination pagination.Pagination `json:"pagination"`
	}
	User struct {
		Id       uint64 `json:"id"`
		Nickname string `json:"nick_name"`
		Username string `json:"user_name"`
		Status   uint8  `json:"status"`
	}
)

func GetUserList(ctx *core.Context, param *GetUserListReq) (*GetUserListRsp, error) {
	var res GetUserListRsp
	users, _, err := admin.GetUserList(ctx, admin.GetUserListParams{
		Page:     param.Page,
		PageSize: param.PageSize,
		Id:       param.Id,
		Username: param.Username,
		Status:   param.Status,
		Mobile:   param.Mobile,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range users {
		res.List = append(res.List, User{
			Id:       item.ID,
			Nickname: item.Nickname,
			Username: item.Username,
			Status:   item.Status,
		})
	}

	// 测试redis
	admin.GetUserListByRedis(ctx)

	// 测试alarm
	ctx.Push("测试alarm", "测试alarm", map[string]interface{}{
		"laya": "sadas",
	})

	ctx.Error("打印error的时候测试一下alarm")

	return &res, nil
}
