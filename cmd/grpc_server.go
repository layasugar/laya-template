package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/middlewares"
	"github.com/layasugar/laya-template/routes"
)

// grpcAppSetup 初始化服务设置
func grpcAppSetup() *laya.App {
	app := laya.GrpcApp()

	// 服务拦截器
	app.GrpcServer().Use(middlewares.TestInterceptor)

	// rpc 路由
	app.GrpcServer().Register(routes.RegisterRpcRoutes)

	return app
}

func main() {
	app := grpcAppSetup()

	app.RunServer()
}
