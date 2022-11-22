package main

import "github.com/layasugar/laya"

// normalAppSetup 初始化基本服务器
func normalAppSetup() *laya.App {
	app := laya.NormalApp()
	return app
}

func main() {
	app := normalAppSetup()

	app.RunServer()
}
