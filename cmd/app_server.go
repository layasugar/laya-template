package main

import "github.com/layasugar/laya"

// defaultAppSetup 初始化基本服务器
func defaultAppSetup() *laya.App {
	app := laya.DefaultApp()

	return app
}

func main() {
	app := webAppSetup()
	//app := defaultAppSetup()
	//app := grpcAppSetup()

	app.RunServer()
}
