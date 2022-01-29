package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/data/test"
)

type (
	Rsp struct {
		Code string `json:"code"`
	}
)

func TraceTest(ctx *laya.WebContext) (*Rsp, error) {
	d, err := test.TraceTest(ctx, test.Req{Status: 1})
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Code,
	}

	return &res, nil
}
