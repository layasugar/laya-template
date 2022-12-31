package admin

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/handle/model/page/admin"
)

// Login 登录 middlewares.UserVerifyToken
func (ctrl *controller) Login(ctx *laya.Context) {
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
func (ctrl *controller) Logout(ctx *laya.Context) {
	err := admin.Logout(ctx)
	if err != nil {
		ctx.Warn("Logout error, err: %s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, "success")
}
