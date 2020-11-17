package laya_go

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"laya-go/middleware"
	"laya-go/ship"
)

func main() {
	// before setting
	ship.Before()

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

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//r.Use(middleware.Sign(), middleware.Response())
	r.Use(middleware.Response())
	service.Handle("/", r)

	// initialise route
	router.Init(r)

	// initialise db
	ship.Init()

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}