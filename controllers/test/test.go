package test

import (
	"context"
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/page/test"
	"github.com/layasugar/laya-template/pb"
)

// HttpTraceTest 测试http请求和链路追踪(http_to_http http_to_grpc)
func (ctrl *controller) HttpTraceTest(ctx *laya.WebContext) {
	// 参数绑定
	var pm test.Req
	err := ctx.ShouldBindJSON(&pm)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	// 参数校验
	if pm.Kind != 1 && pm.Kind != 2 {
		ctrl.Fail(ctx, errors.New("kind 只能是1,2"))
	}

	// 业务处理
	res, err := test.HttpTraceTest(ctx, pm)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	// 响应
	ctrl.Suc(ctx, res)
}

func (ctrl *controller) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Name}, nil
}

// GrpcTraceTest 测试http请求和链路追踪(grpc_to_http grpc_to_grpc)
func (ctrl *controller) GrpcTraceTest(ctx context.Context, in *pb.GrpcTraceTestReq) (*pb.HelloReply, error) {
	// 转换ctx
	newCtx := ctx.(*laya.GrpcContext)

	// 参数验证
	if in.Kind == 0 {
		return nil, errors.New("请传入kind")
	}

	// 业务处理
	resp, err := test.RpcTraceTest(newCtx, in)
	if err != nil {
		return nil, err
	}

	// 响应
	return &pb.HelloReply{Message: resp.Code}, nil
}
