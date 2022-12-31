package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/handle/middleware"
)

// webAppSetup 初始化服务设置
func webAppSetup() *laya.App {
	app := laya.WebApp()

	// register global middlewares
	app.WebServer().Use(middleware.TestMiddleware())

	// register routes
	app.WebServer().Register(routes.Register)

	return app
}

func main() {
	app := webAppSetup()

	app.RunServer()
}
