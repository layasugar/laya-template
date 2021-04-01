package main

import (
	"github.com/layatips/laya"
	"github.com/layatips/laya-go/routes"
	"github.com/layatips/laya/opts"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.NewApp()
	app.WebServer().RegisterRouter(routes.RoutingSetup)
	app.Use(opts.Db, opts.Mdb, opts.Rdb, opts.Mem)
	return app
}

func main() {
	app := ServerSetup()
	app.RunWebServer()
}
