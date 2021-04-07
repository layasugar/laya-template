package main

import (
	"github.com/layatips/laya"
	"github.com/layatips/laya-go/models/dao"
	"github.com/layatips/laya-go/routes"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.NewApp()
	app.RegisterRouter(routes.Routes)
	app.Use(dao.InitDao)
	return app
}

func main() {
	app := ServerSetup()
	app.RunWebServer()
}
