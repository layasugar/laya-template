package handler

import (
	"fmt"

	"github.com/layasugar/laya"
	"github.com/layasugar/laya/gcnf"

	"github.com/layasugar/laya-template/global"
)

// Ctrl the controllers with some useful and common function
var _ = &BaseCtrl{}

type BaseCtrl struct {
	*global.HttpResp
}

// Version version
func (ctrl *BaseCtrl) Version(ctx *laya.Context) {
	res := fmt.Sprintf("%s version: %s\nlisten: %s", gcnf.AppName(), gcnf.AppVersion(), gcnf.Listen())
	ctx.Info("测试日志%s", "hello world")
	_, _ = ctx.Gin().Writer.Write([]byte(res))
}

func (ctrl *BaseCtrl) Ready(ctx *laya.Context) {
	_, _ = ctx.Gin().Writer.Write([]byte("ok, I'm fine"))
}

func (ctrl *BaseCtrl) Healthy(ctx *laya.Context) {

}
