package main

import (
	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/middleware"
	"github.com/layasugar/laya-template/route"
)

// grpcAppSetup 初始化服务设置
func grpcAppSetup() *laya.App {
	app := laya.GrpcApp()

	// 服务拦截器
	app.GrpcServer().Use(middleware.TestInterceptor)

	// rpc 路由
	app.GrpcServer().Register(route.RegisterRpcRoutes)

	return app
}

func main() {
	app := grpcAppSetup()

	app.RunServer()
}
