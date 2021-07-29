package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya-go/global/errno"
	"github.com/layasugar/laya-go/models/data"
	"github.com/layasugar/laya/glogs"
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
		glogs.WarnF(c, "用户中心", "用户不存在")
		return nil, errno.UserNotFound
	}
	return user, nil
}
