package handler

import "github.com/layasugar/laya-template/pkg/core"

// Test test接口
func (ctrl *BaseCtrl) Test(ctx *core.Context) {
	var body map[string]interface{}
	_ = ctx.Gin().ShouldBindJSON(&body)
	ctrl.Suc(ctx, body, "this is http_test restfulApi")
}
