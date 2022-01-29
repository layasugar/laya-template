package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao/http_cal/test"
)

type (
	Req struct {
		Status int64 `json:"status"`
	}
	Rsp struct {
		Code string `json:"code"`
	}
)

func TraceTest(ctx *laya.WebContext, pm Req) (*Rsp, error) {
	d, err := test.GetTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Code,
	}

	return &res, nil
}
