package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/test"
)

func (ctrl *controller) TraceTest(ctx *laya.WebContext) {
	res, err := test.TraceTest(ctx)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	ctrl.Suc(ctx, res)
}
