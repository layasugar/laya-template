package app

import (
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/app"
)

// Login 登录
func (ctrl *controller) Login(ctx *laya.WebContext) {
	// 参数绑定
	var pm app.LoginParam
	if err := ctx.ShouldBindJSON(&pm); err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	// 参数校验
	if is := tools.IsMobile(pm.Mobile); !is {
		ctrl.Fail(ctx, errors.New("请输入正确的手机号码"))
		return
	}

	// 业务处理
	data, err := app.Login(ctx, &pm)
	if err != nil {
		ctx.ErrorF("Login error, err: %s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}

	// 响应数据
	ctrl.Suc(ctx, data)
}

// Logout 退出登录
func (ctrl *controller) Logout(ctx *laya.WebContext) {
	err := app.Logout(ctx)
	if err != nil {
		ctx.WarnF("Logout error, err: %s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, "success")
}