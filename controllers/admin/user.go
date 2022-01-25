package admin

import (
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/admin"
)

type getUserListParam = admin.UserParam

func (ctrl *controller) GetUserList(ctx *laya.WebContext) {
	var param getUserListParam
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	resp, err := admin.GetUserList(ctx, &param)
	if err != nil {
		ctx.Info("获取用户列表", fmt.Sprintf("title=获取用户列表,err=%s", err.Error()))
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, resp)
}
