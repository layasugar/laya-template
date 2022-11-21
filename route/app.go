package route

import (
	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/handle/app"
	"github.com/layasugar/laya-template/middleware"
)

func RegisterApp(r *laya.WebServer) {
	appGroupV1 := r.Group("/app/v1")
	{
		appGroupV1.POST("/login", app.Ctrl.Login)
		appGroupV1.POST("/logout", app.Ctrl.Logout).Use(middleware.UserVerifyToken())
		appGroupV1.POST("/user/info", app.Ctrl.GetUserInfo).Use(middleware.UserVerifyToken())
	}
}
