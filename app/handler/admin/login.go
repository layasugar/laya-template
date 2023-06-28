package admin

import (
	"github.com/layasugar/laya-template/app/models/page/admin"
	"github.com/layasugar/laya-template/pkg/core"
)

// Login 登录 middlewares.UserVerifyToken
func (ctrl *controller) Login(ctx *core.Context) {
	var pm admin.LoginReq
	if err := ctx.Gin().ShouldBind(&pm); err != nil {
		ctrl.Fail(ctx, err)
		return
	}
	data, err := admin.Login(ctx, &pm)
	if err != nil {
		ctx.Warn("Login error,err:%s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, data)
}

// Logout 退出登录
func (ctrl *controller) Logout(ctx *core.Context) {
	err := admin.Logout(ctx)
	if err != nil {
		ctx.Warn("Logout error, err: %s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, "success")
}
