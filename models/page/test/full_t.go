package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/data/test"
	"github.com/layasugar/laya-template/pb"
)

type (
	Req struct {
		Kind uint8 `json:"kind"`
	}

	Rsp struct {
		Code string `json:"code"`
	}
)

func FullTest(ctx *laya.WebContext, pm Req) (*Rsp, error) {
	var res Rsp
	switch pm.Kind {
	case 1:
		res.Code = "success"
	case 2:
		mysqlTestCurd(ctx)
		res.Code = "success"
	case 3:
		test.RedisTestCurd(ctx)
		res.Code = "success"
	case 4:
		mongoTestCurd(ctx)
		res.Code = "success"
	case 5:
		esTestCurd(ctx)
		res.Code = "success"
	case 6:
		d, err := test.HttpToHttpTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	case 7:
		d, err := test.HttpToRpcTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	}

	return &res, nil
}

func RpcTraceTest(ctx *laya.GrpcContext, pm *pb.GrpcTraceTestReq) (*Rsp, error) {
	var res Rsp
	switch pm.Kind {
	case 1:
		d, err := test.RpcToHttpTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	case 2:
		d, err := test.RpcToRpcTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	}

	return &res, nil
}
