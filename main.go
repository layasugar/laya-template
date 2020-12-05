package main

import (
	"github.com/LaYa-op/laya"
	"github.com/LaYa-op/laya-go/routes"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.NewApp()
	app.WebServer().RegisterRouter(routes.RoutingSetup)


	// 初始化内存缓存
	//memory.Init()
	return app
}

func main() {
	app := ServerSetup()
	app.RunWebServer()
}
