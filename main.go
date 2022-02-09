package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/routes"
	"runtime"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.DefaultApp()

	// open db connection and global do before something
	app.Use(dao.Init)

	// register global middlewares
	//app.WebServer().Use()

	// register routes
	app.WebServer().RegisterRouter(routes.Register)

	// rpc 路由
	//routes.RegisterRpcRoutes(app.PbRPCServer())

	// 屏蔽不需要打印出入参路由分组
	app.SetNoLogParamsPrefix("/admin")

	return app
}

func main() {
	app := ServerSetup()

	runtime.Gosched()
	app.RunWebServer()
	//app.RunPbRPCServer()
}
