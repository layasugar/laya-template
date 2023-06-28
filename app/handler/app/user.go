package app

import (
	"fmt"

	"github.com/layasugar/laya-template/app/models/page/app"
	"github.com/layasugar/laya-template/pkg/core"
)

// GetUserInfo 获取用户信息
func (ctrl *controller) GetUserInfo(ctx *core.Context) {
	resp, err := app.GetUserInfo(ctx)
	if err != nil {
		ctx.Info("获取用户信息", fmt.Sprintf("title=获取用户信息,err=%s", err.Error()))
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, resp)
}
