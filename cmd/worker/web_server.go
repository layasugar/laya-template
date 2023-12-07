package worker

import (
	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/routes"
	"github.com/layasugar/laya-template/routes/middlewares"
	"testing"
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

func Web(t *testing.T) {
	app := webAppSetup()

	app.RunServer()
}
