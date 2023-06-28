package core

import (
	"context"
	"log"
	"net"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/core/metautils"
	"github.com/layasugar/laya/core/util"
	"github.com/layasugar/laya/gcnf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GrpcServer struct
type GrpcServer struct {
	*grpc.Server
	opts   []grpc.UnaryServerInterceptor
	routes []func(server *GrpcServer)
}

// newGrpcServer create new GrpcServer with default configuration
func newGrpcServer() *GrpcServer {
	server := &GrpcServer{
		opts: []grpc.UnaryServerInterceptor{
			serverInterceptor,
		},
	}

	return server
}

func (gs *GrpcServer) Use(f ...func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)) {
	for _, vf := range f {
		gs.opts = append(gs.opts, vf)
	}
}

func (gs *GrpcServer) Register(f ...func(s *GrpcServer)) {
	gs.routes = append(gs.routes, f...)
}

func (gs *GrpcServer) Run(addr string) (err error) {
	// 初始化server, 将多个拦截器构建成一个拦截器
	gs.Server = grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(gs.opts...)),
	)

	// 注册路由
	for _, vf := range gs.routes {
		vf(gs)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}

	// 在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(gs.Server)

	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们
	err = gs.Server.Serve(lis)
	return
}

// serverInterceptor 提供服务的拦截器, 重写context, 记录出入参, 记录链路追踪
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md := metautils.ExtractIncoming(ctx)
	newCtx := NewContext(constants.SERVERGRPC, info.FullMethod, md, nil)
	reqByte, _ := util.CJson.Marshal(req)
	mdByte, _ := util.CJson.Marshal(md)
	resp, err := handler(newCtx, req)
	respByte, _ := util.CJson.Marshal(resp)
	if gcnf.CheckLogParams(info.FullMethod) {
		newCtx.Info("params_log", newCtx.Field("header", string(mdByte)),
			newCtx.Field("path", info.FullMethod), newCtx.Field("protocol", constants.PROTOCOLGRPC),
			newCtx.Field("inbound", string(reqByte)), newCtx.Field("outbound", string(respByte)))
	}
	newCtx.End(newCtx.TopSpan())
	return resp, err
}
