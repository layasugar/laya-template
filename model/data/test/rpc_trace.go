package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/model/dao/cal/rpc_test"
)

func RpcToHttpTraceTest(ctx *laya.Context) (*Rsp, error) {
	d, err := rpc_test.HttpTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Code,
	}

	return &res, nil
}

func RpcToRpcTraceTest(ctx *laya.Context) (*Rsp, error) {
	d, err := rpc_test.RpcTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Message,
	}

	return &res, nil
}
