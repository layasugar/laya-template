package admin

import (
	"github.com/layasugar/laya-template/app/models/dao"
	"github.com/layasugar/laya-template/app/models/dao/db"
	"github.com/layasugar/laya-template/pkg/core"
	"log"
)

type (
	User              = db.User
	GetUserListParams struct {
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
		Id       uint64 `json:"id"`
		Username string `json:"username"`
		Status   uint8  `json:"status"`
		Mobile   string `json:"mobile"`
	}
)

func GetUserList(ctx *core.Context, params GetUserListParams) ([]User, int64, error) {
	var users []User
	tx := dao.Orm(ctx)
	if params.Id > 0 {
		tx.Where("id = ?", params.Id)
	}
	if params.Username != "" {
		tx.Where("username = ?", params.Username)
	}
	if params.Status > 0 {
		tx.Where("status = ?", params.Status)
	}
	if params.Mobile != "" {
		tx.Where("mobile = ?", params.Mobile)
	}

	var total int64
	err := tx.Model(&users).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if params.Page > 0 && params.PageSize > 0 {
		tx.Limit(params.PageSize).Offset(params.Page * params.PageSize)
	}

	err = tx.Find(&users).Error
	return users, total, err
}

func GetUserListByRedis(ctx *core.Context) {
	result, err := dao.Rdb().Get(ctx, "aksda").Result()
	if err != nil {
		ctx.Info("asdasd, err: %s", err.Error())
	}

	log.Print(result)
}
