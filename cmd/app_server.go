package main

import (
	"testing"

	"github.com/layasugar/laya-template/pkg/core"
)

// normalAppSetup 初始化基本服务器
func normalAppSetup() *core.App {
	app := core.NormalApp()
	return app
}

func App(t *testing.T) {
	app := normalAppSetup()

	app.RunServer()
}
