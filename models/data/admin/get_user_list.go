package admin

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/models/dao/db"
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

func GetUserList(ctx *laya.WebContext, params GetUserListParams) ([]User, int64, error) {
	var users []User
	tx := dao.DB.WithContext(ctx)
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
	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if params.Page > 0 && params.PageSize > 0 {
		tx.Limit(params.PageSize).Offset(params.Page * params.PageSize)
	}

	err = tx.Find(&users).Error
	return users, total, err
}
