package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/global/errno"
	"github.com/layatips/laya-go/models/data"
	"gorm.io/gorm"
)

type UserParam struct {
	Id uint64 `json:"id" binding:"required"`
}

func GetUserInfo(c *gin.Context, param *UserParam) (interface{}, error) {
	user, err := data.GetUserById(c, param.Id)
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errno.UserNotFound
	}
	return user, nil
}
