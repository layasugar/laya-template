package test

import (
	"github.com/layasugar/laya-template/app/models/dao/cal/rpc_test"
	"github.com/layasugar/laya-template/pkg/core"
)

func RpcToHttpTraceTest(ctx *core.Context) (*Rsp, error) {
	d, err := rpc_test.HttpTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Code,
	}

	return &res, nil
}

func RpcToRpcTraceTest(ctx *core.Context) (*Rsp, error) {
	d, err := rpc_test.RpcTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Message,
	}

	return &res, nil
}
