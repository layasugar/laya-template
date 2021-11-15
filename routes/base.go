package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya-go/controllers"
)

func Register(r *gin.Engine) {
	r.GET("/", controllers.Ctrl.Version)   //version
	r.POST("/test", controllers.Ctrl.Test) //测试接口
	RegisterApp(r)
}
