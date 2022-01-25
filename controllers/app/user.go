package app

import (
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/app"
)

type getUserInfoParam = app.UserParam

// GetUserInfo 获取用户信息
func (ctrl *BaseAppCtrl) GetUserInfo(ctx *laya.WebContext) {
	var param getUserInfoParam
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	resp, err := app.GetUserInfo(ctx, &param)
	if err != nil {
		ctx.Info("获取用户信息", fmt.Sprintf("title=获取用户信息,err=%s", err.Error()))
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, resp)
}
