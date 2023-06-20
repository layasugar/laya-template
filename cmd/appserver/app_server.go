package main

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global/version"
)

// normalAppSetup 初始化基本服务器
func normalAppSetup() *laya.App {
	app := laya.NormalApp()
	return app
}

func main() {
	app := normalAppSetup()

	version.Birth
	app.RunServer()
}
