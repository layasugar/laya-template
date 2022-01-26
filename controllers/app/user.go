package app

import (
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/app"
)

// GetUserInfo 获取用户信息
func (ctrl *controller) GetUserInfo(ctx *laya.WebContext) {
	resp, err := app.GetUserInfo(ctx)
	if err != nil {
		ctx.Infof("获取用户信息", fmt.Sprintf("title=获取用户信息,err=%s", err.Error()))
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, resp)
}
