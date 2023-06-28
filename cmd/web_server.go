package main

import (
	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/routes"
	"github.com/layasugar/laya-template/routes/middlewares"
)

// webAppSetup 初始化服务设置
func webAppSetup() *core.App {
	app := core.WebApp()

	// register global middlewares
	app.WebServer().Use(middlewares.TestMiddleware())

	// register routes
	app.WebServer().Register(routes.Register)

	return app
}

func main() {
	app := webAppSetup()

	app.RunServer()
}
