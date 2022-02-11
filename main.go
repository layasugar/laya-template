package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/routes"
)

// webAppSetup 初始化服务设置
func webAppSetup() *laya.App {
	app := laya.WebApp()

	// open db connection and global do before something
	app.Use(dao.Init)

	// register global middlewares
	//app.WebServer().Use()

	// register routes
	app.WebServer().Register(routes.Register)

	// 屏蔽不需要打印出入参路由分组
	app.SetNoLogParamsPrefix("/admin")

	return app
}

// grpcAppSetup 初始化服务设置
func grpcAppSetup() *laya.App {
	app := laya.GrpcApp()

	// open db connection and global do before something
	app.Use(dao.Init)

	// rpc 路由
	app.GrpcServer().Register(routes.RegisterRpcRoutes)

	return app
}

func main() {
	//app := webAppSetup()
	app := grpcAppSetup()

	app.RunServer()
}
