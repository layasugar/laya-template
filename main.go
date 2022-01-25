package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/middleware"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/routes"
)

// ServerSetup 初始化服务设置
func ServerSetup() *laya.App {
	app := laya.DefaultApp()

	// open db connection and global do before something
	app.Use(dao.Init)

	// register middleware
	app.WebServer().Use(middleware.SetHeader, middleware.LogInParams, middleware.SetTrace)

	// register routes
	app.WebServer().RegisterRouter(routes.Register)

	// rpc 路由
	//app.PbRPCServer().AddHandler(rpc.AddUser)

	// 屏蔽不需要打印出入参路由分组
	global.SetNoLogParamsPrefix("/admin")

	return app
}

func main() {
	app := ServerSetup()

	app.RunWebServer()
	//app.RunPbRPCServer()
}
