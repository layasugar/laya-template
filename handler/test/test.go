package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/layasugar/laya-template/handle/model/page/test"
	"github.com/layasugar/laya-template/handle/pb"

	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/utils"
)

// FullTest 测试http请求和链路追踪(http_to_http http_to_grpc)
func (ctrl *controller) FullTest(ctx *laya.Context) {
	// 参数绑定
	var pm test.Req
	err := ctx.Gin().ShouldBindJSON(&pm)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	// 参数校验
	var kinds = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !utils.InSliceUint8(pm.Kind, kinds) {
		ctrl.Fail(ctx, fmt.Errorf("kind 只能是%v", kinds))
		return
	}

	// 业务处理
	res, err := test.FullTest(ctx, pm)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	// 响应
	ctrl.Suc(ctx, res)
}

// SayHello 基础测试使用
func (ctrl *controller) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Name}, nil
}

// GrpcTraceTest 测试http请求和链路追踪(grpc_to_http grpc_to_grpc)
func (ctrl *controller) GrpcTraceTest(ctx context.Context, in *pb.GrpcTraceTestReq) (*pb.HelloReply, error) {
	// 转换ctx
	newCtx := ctx.(*laya.Context)

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
