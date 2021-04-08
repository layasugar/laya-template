package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/global/errno"
	"github.com/layatips/laya-go/models/data"
	"github.com/layatips/laya/glogs"
	"gorm.io/gorm"
)

type UserParam struct {
	Id uint64 `json:"id" binding:"required"`
}

func GetUserInfo(c *gin.Context, param *UserParam) (interface{}, error) {
	user, err := data.GetUserById(c, param.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		glogs.WarnFR(c, "%s", "用户不存在")
		return nil, errno.UserNotFound
	}
	return user, nil
}
