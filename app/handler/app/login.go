package app

import (
	"errors"
	"github.com/layasugar/laya-template/app/models/page/app"
	"github.com/layasugar/laya-template/pkg/core"

	"github.com/layasugar/laya-template/utils"
)

// Login 登录
func (ctrl *controller) Login(ctx *core.Context) {
	// 参数绑定
	var pm app.LoginParam
	if err := ctx.Gin().ShouldBindJSON(&pm); err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	// 参数校验
	if is := utils.IsMobile(pm.Mobile); !is {
		ctrl.Fail(ctx, errors.New("请输入正确的手机号码"))
		return
	}

	// 业务处理
	data, err := app.Login(ctx, &pm)
	if err != nil {
		ctx.Error("Login error, err: %s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}

	// 响应数据
	ctrl.Suc(ctx, data)
}

// Logout 退出登录
func (ctrl *controller) Logout(ctx *core.Context) {
	err := app.Logout(ctx)
	if err != nil {
		ctx.Warn("Logout error, err: %s ", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, "success")
}
