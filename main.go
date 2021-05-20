package main

import (
	"github.com/layatips/laya"
	"github.com/layatips/laya-go/middleware"
	"github.com/layatips/laya-go/routes"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.DefaultApp()

	// open file watcher
	//app.RegisterFileWatcher(genv.ConfigPath, global.ConfChangeHandler)

	// open db connection and global do before something
	//app.Use(dao.Init, global.Init)

	// register middleware
	app.WebServer.Use(middleware.SetHeader, middleware.LogInParams, middleware.SetTrace)

	// register routes
	app.RegisterRouter(routes.Register)

	return app
}

func main() {
	app := ServerSetup()
	app.RunWebServer()
}
