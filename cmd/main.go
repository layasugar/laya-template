package main

import (
	"fmt"
	"os"

	"github.com/layasugar/laya-template/routes"
	"github.com/layasugar/laya-template/routes/middlewares"
	"github.com/layasugar/laya-template/utils/core"
)

// appSetup 初始化服务设置
func appSetup() *core.App {
	app := core.WebApp()

	// register global middlewares
	app.WebServer().Use(middlewares.TestMiddleware())

	// register routes
	app.WebServer().Register(routes.Register)

	return app
}

const pidFile = "/var/run/layatp.pid"

func main() {
	// 启动后获取当前进程的pid
	pid := os.Getpid()
	// 将pid写入文件
	err := os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", pid)), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", pid)

	app := appSetup()

	app.RunServer()
}
