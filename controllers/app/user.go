package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya-go/models/page/app"
	"github.com/layasugar/glogs"
)

type getUserInfoParam = app.UserParam

// 获取用户信息
func (ctrl *BaseAppCtrl) GetUserInfo(c *gin.Context) {
	var param getUserInfoParam
	err := c.ShouldBind(&param)
	if err != nil {
		ctrl.Fail(c, err)
		return
	}

	resp, err := app.GetUserInfo(c, &param)
	if err != nil {
		glogs.ErrorF(c, "获取用户信息", fmt.Sprintf("title=获取用户信息,err=%s", err.Error()))
		ctrl.Fail(c, err)
		return
	}
	ctrl.Suc(c, resp)
}
