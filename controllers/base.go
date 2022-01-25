package controllers

import (
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya/genv"
)

// Ctrl the controllers with some useful and common function
var Ctrl = &BaseCtrl{}

type BaseCtrl struct {
	global.Resp
}

// Version version
func (ctrl *BaseCtrl) Version(ctx *laya.WebContext) {
	res := fmt.Sprintf("%s version: %s\napp_url: %s", genv.AppName(), genv.AppVersion(), genv.AppUrl())
	ctx.Info("测试日志%s", "hello world")
	_, _ = ctx.Writer.Write([]byte(res))
	return
}
