package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/controllers"
)

func Register(r *laya.WebServer) {
	r.GET("/", controllers.Ctrl.Version)   // version
	r.POST("/test", controllers.Ctrl.Test) // 测试接口
	RegisterApp(r)
}
