package app

import (
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/models/page/app"
	"github.com/layatips/laya/glogs"
)

type getUserInfoParam = app.UserParam

// 获取用户信息
func (ctrl *controller) GetUserInfo(c *gin.Context) {
	var param getUserInfoParam
	err := c.ShouldBind(&param)
	if err != nil {
		ctrl.Fail(c, err)
		return
	}

	resp, err := app.GetUserInfo(c, &param)
	if err != nil {
		glogs.ErrorFR(c, "title=获取用户信息,err=%s", err.Error())
		ctrl.Fail(c, err)
		return
	}
	ctrl.Suc(c, resp)
}
