package admin

import (
	"fmt"
	"github.com/layasugar/glogs"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/admin"
)

type getUserListParam = admin.UserParam

func (ctrl *controller) GetUserList(c *laya.WebContext) {
	var param getUserListParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		ctrl.Fail(c, err)
		return
	}

	resp, err := admin.GetUserList(c, &param)
	if err != nil {
		glogs.ErrorF(c.Request, "获取用户列表", fmt.Sprintf("title=获取用户列表,err=%s", err.Error()))
		ctrl.Fail(c, err)
		return
	}
	ctrl.Suc(c, resp)
}
