package main

import (
	"laya-go/server/hall/middleware"
	"laya-go/server/hall/router"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"io"
	"os"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("tb.server.hall"),
		web.Version("1.0"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	_ = os.MkdirAll("logs", 777)
	f, _ := os.Create("logs/hallServer.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(middleware.Sign())
	service.Handle("/", r)

	// main route
	router.Init(r)
	// main db
	Init()

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
