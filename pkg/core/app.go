// surprise

package core

import (
	"flag"
	"fmt"
	"log"

	"github.com/layasugar/laya-template/pkg/core/constants"
	"github.com/layasugar/laya-template/pkg/core/logger"
	"github.com/layasugar/laya-template/pkg/gcnf"
)

type App struct {
	serverType constants.SERVERTYPE // serverType
	webServer  *WebServer           // webServer 目前web引擎使用gin
	grpcServer *GrpcServer          // grpcServer
}

// NormalApp 默认应用不带有web或者grpc, 可作为服务使用
func NormalApp() *App {
	app := new(App)
	app.serverType = constants.SERVERNORMAL
	app.initWithConfig()
	return app
}

// WebApp web app
func WebApp() *App {
	app := new(App)
	app.serverType = constants.SERVERGIN
	app.initWithConfig()
	return app
}

// GrpcApp grpc app
func GrpcApp() *App {
	app := new(App)
	app.serverType = constants.SERVERGRPC
	app.initWithConfig()
	return app
}

// 初始化app
func (app *App) initWithConfig() *App {
	// 接收命令行参数
	var f string
	flag.StringVar(&f, "config", "", "set a config file")
	flag.Parse()

	// 初始化配置
	err := gcnf.InitConfig(f)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	logger.InitSugar(&logger.Config{
		AppName:       gcnf.AppName(),
		AppMode:       gcnf.AppMode(),
		ChildPath:     gcnf.LogChildPath(),
		LogPath:       gcnf.LogPath(),
		LogType:       gcnf.LogType(),
		LogLevel:      gcnf.LogLevel(),
		RotationSize:  gcnf.LogMaxSize(),
		RotationCount: gcnf.LogMaxCount(),
		RotationTime:  gcnf.LogMaxTime(),
		MaxAge:        gcnf.LogMaxAge(),
	})

	switch app.serverType {
	case constants.SERVERGIN:
		app.webServer = newWebServer(gcnf.GinRunMode())
		if len(defaultWebServerMiddlewares) > 0 {
			app.webServer.Use(defaultWebServerMiddlewares...)
		}
	case constants.SERVERGRPC:
		app.grpcServer = newGrpcServer()
	}
	return app
}

// RunServer 运行Web服务
func (app *App) RunServer() {
	switch app.serverType {
	case constants.SERVERGIN:
		// 启动web服务
		log.Printf("[app] Listening and serving %s on %s\n", "HTTP", gcnf.Listen())
		err := app.webServer.Run(gcnf.Listen())
		if err != nil {
			fmt.Printf("Can't RunWebServer: %v\n", err)
		}
	case constants.SERVERGRPC:
		// 启动grpc服务
		log.Printf("[app] Listening and serving %s on %s\n", "GRPC", gcnf.Listen())
		err := app.grpcServer.Run(gcnf.Listen())
		if err != nil {
			log.Fatalf("Can't RunGrpcServer, GrpcListen: %s, err: %v", gcnf.Listen(), err)
		}
	case constants.SERVERNORMAL:
	}
}

// Use 提供一个加载函数
func (app *App) Use(fc ...func()) {
	for _, f := range fc {
		f()
	}
}

// WebServer 获取WebServer的指针
func (app *App) WebServer() *WebServer {
	return app.webServer
}

// GrpcServer 获取PbRPCServer的指针
func (app *App) GrpcServer() *GrpcServer {
	return app.grpcServer
}
