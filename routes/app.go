package routes

import (
	"github.com/layasugar/laya-template/app/handler/app"
	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/routes/middlewares"
)

func RegisterApp(r *core.WebServer) {
	appGroupV1 := r.Group("/app/v1")
	{
		appGroupV1.POST("/login", app.Ctrl.Login)
		appGroupV1.POST("/logout", app.Ctrl.Logout).Use(middlewares.UserVerifyToken())
		appGroupV1.POST("/user/info", app.Ctrl.GetUserInfo).Use(middlewares.UserVerifyToken())
	}
}
