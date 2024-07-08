package main

import (
	"testing"

	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/routes"
	"github.com/layasugar/laya-template/routes/middlewares"
)

// grpcAppSetup 初始化服务设置
func grpcAppSetup() *core.App {
	app := core.GrpcApp()

	// 服务拦截器
	app.GrpcServer().Use(middlewares.TestInterceptor)

	// rpc 路由
	app.GrpcServer().Register(routes.RegisterRpcRoutes)

	return app
}

func Grpc(t *testing.T) {
	app := grpcAppSetup()

	app.RunServer()
}
