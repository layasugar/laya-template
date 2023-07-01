package test

import (
	"github.com/layasugar/laya-template/app/models/dao/cal/http_test"
	"github.com/layasugar/laya-template/pkg/core"
)

func HttpToHttpTraceTest(ctx *core.Context) (*Rsp, error) {
	d, err := http_test.HttpToHttpTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Code,
	}

	return &res, nil
}

func HttpToRpcTraceTest(ctx *core.Context) (*Rsp, error) {
	d, err := http_test.HttpToGrpcTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Message,
	}

	return &res, nil
}
