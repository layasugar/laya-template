package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"laya-go/server/hall/router"
	"laya-go/ship"
	"laya-go/ship/middleware"
)

func main() {
	// before setting
	ship.Before()

	// create new web service
	service := web.NewService(
		web.Name("tb.server.hall"),
		web.Version("1.0"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.Use(middleware.Sign(), middleware.Response())
	r.Use(middleware.Response())
	service.Handle("/", r)

	// initialise route
	router.Init(r)

	// initialise db
	ship.Init(RC, MC)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
