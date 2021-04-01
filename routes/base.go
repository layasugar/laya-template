package routes

import (
	"github.com/layatips/laya"
	"github.com/layatips/laya-go/controllers"
)

func RoutingSetup(r *laya.WebServer) {
	r.GET("/", controllers.Ctrl.Version)                    //version
	r.POST("/test", controllers.Ctrl.Test)                  //测试接口
	r.POST("/memory-status", controllers.Ctrl.MemoryStatus) //测试接口
	r.GET("/health", controllers.Ctrl.HealthCheck)          //存活探针
	r.GET("/ready", controllers.Ctrl.ReadyCheck)            //业务探针
	r.GET("/reload", controllers.Ctrl.Reload)               //配置重载
	RegisterRoute(r)
}
