package controllers

import (
	"github.com/layasugar/laya"
)

// Test test接口
func (ctrl *BaseCtrl) Test(ctx *laya.WebContext) {
	var body map[string]interface{}
	_ = ctx.ShouldBindJSON(&body)
	ctrl.Suc(ctx, body, "this is http_test restfulApi")
}
