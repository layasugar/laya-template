package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/models/data"
)

type UserParam struct {
	Zone string `json:"zone" binding:"required"`
}

func GetUserList(c *gin.Context, param *UserParam) (interface{}, error) {
	users, err := data.GetUserListByZone(c, param.Zone)
	return users, err
}
