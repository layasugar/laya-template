package main

import (
	"github.com/layatips/laya"
	"github.com/layatips/laya-go/models/dao"
	"github.com/layatips/laya-go/routes"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.NewApp()
	app.Use(dao.Init)
	// register middleware
	//app.WebServer.Use(middleware.LogInParams)
	app.RegisterRouter(routes.Register)
	return app
}

func main() {
	app := ServerSetup()
	app.RunWebServer()
}
