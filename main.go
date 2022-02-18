package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/middlewares"
	"github.com/layasugar/laya-template/routes"
)

// webAppSetup 初始化服务设置
func webAppSetup() *laya.App {
	app := laya.WebApp()

	// register global middlewares
	app.WebServer().Use(middlewares.TestMiddleware())

	// register routes
	app.WebServer().Register(routes.Register)

	// 屏蔽不需要打印出入参路由分组
	app.SetNoLogParamsPrefix("/admin")

	return app
}

// grpcAppSetup 初始化服务设置
func grpcAppSetup() *laya.App {
	app := laya.GrpcApp()

	// 服务拦截器
	app.GrpcServer().Use(middlewares.TestInterceptor)

	// rpc 路由
	app.GrpcServer().Register(routes.RegisterRpcRoutes)

	return app
}

// defaultAppSetup 初始化基本服务器
func defaultAppSetup() *laya.App {
	app := laya.DefaultApp()

	return app
}

func main() {
	app := webAppSetup()
	//app := defaultAppSetup()
	//app := grpcAppSetup()

	app.RunServer()
}
