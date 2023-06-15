package handler

import (
	"github.com/layasugar/laya"
)

// Test test接口
func (ctrl *controllers.BaseCtrl) Test(ctx *laya.Context) {
	var body map[string]interface{}
	_ = ctx.Gin().ShouldBindJSON(&body)
	ctrl.Suc(ctx, body, "this is http_test restfulApi")
}
