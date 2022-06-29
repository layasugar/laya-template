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

func main() {
	app := webAppSetup()
	//app := defaultAppSetup()
	//app := grpcAppSetup()

	app.RunServer()
}
