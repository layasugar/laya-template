package app

import (
	"fmt"
	"github.com/layasugar/glogs"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/app"
)

type getUserInfoParam = app.UserParam

// GetUserInfo 获取用户信息
func (ctrl *BaseAppCtrl) GetUserInfo(c *laya.WebContext) {
	var param getUserInfoParam
	err := c.ShouldBind(&param)
	if err != nil {
		ctrl.Fail(c, err)
		return
	}

	resp, err := app.GetUserInfo(c, &param)
	if err != nil {
		glogs.ErrorF(c.Request, "获取用户信息", fmt.Sprintf("title=获取用户信息,err=%s", err.Error()))
		ctrl.Fail(c, err)
		return
	}
	ctrl.Suc(c, resp)
}
