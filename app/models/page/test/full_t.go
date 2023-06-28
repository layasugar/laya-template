package test

import (
	test2 "github.com/layasugar/laya-template/handle/model/data/test"
	"github.com/layasugar/laya-template/handle/pb"
)

type (
	Req struct {
		Kind uint8 `json:"kind"`
	}

	Rsp struct {
		Code string `json:"code"`
	}
)

func FullTest(ctx *core.Context, pm Req) (*Rsp, error) {
	var res Rsp
	switch pm.Kind {
	case 1:
		res.Code = "success"
	case 2:
		mysqlTestCurd(ctx)
		res.Code = "success"
	case 3:
		test2.RedisTestCurd(ctx)
		res.Code = "success"
	case 4:
		mongoTestCurd(ctx)
		res.Code = "success"
	case 5:
		esTestCurd(ctx)
		res.Code = "success"
	case 6:
		d, err := test2.HttpToHttpTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	case 7:
		d, err := test2.HttpToRpcTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	}

	return &res, nil
}

func RpcTraceTest(ctx *core.Context, pm *pb.GrpcTraceTestReq) (*Rsp, error) {
	var res Rsp
	switch pm.Kind {
	case 1:
		d, err := test2.RpcToHttpTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	case 2:
		d, err := test2.RpcToRpcTraceTest(ctx)
		if err != nil {
			return nil, err
		}

		res.Code = d.Code
	}

	return &res, nil
}
