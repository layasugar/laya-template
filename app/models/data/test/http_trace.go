package test

import (
	"github.com/layasugar/laya-template/handle/model/dao/cal/http_test"
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
