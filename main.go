package main

import (
	"github.com/LaYa-op/laya"
	"github.com/LaYa-op/laya-go/middleware"
	"github.com/LaYa-op/laya-go/routers"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("tb.controllers.hall"),
		web.Version("1.0"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	//r.Use(middleware.Sign(), middleware.Response())
	//r.Use(middleware.Base.Auth)
	r.Use(middleware.Base.Response)
	//r.Use(middleware.Base.Sign)
	//r.Use(middleware.Base.Test)
	service.Handle("/", r)

	// initialise route
	routers.Init(r)

	// initialise db
	laya.Init()

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// before setting
	laya.Before()
	log.Info("init the server")
}
