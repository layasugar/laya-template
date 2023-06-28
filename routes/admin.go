package routes

import (
	"github.com/layasugar/laya-template/app/handler/admin"
	"github.com/layasugar/laya-template/app/handler/app"
	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/routes/middlewares"
)

func RegisterAdmin(r *core.WebServer) {
	adminGroup := r.Group("/admin/v1")
	{
		adminGroup.POST("/login", admin.Ctrl.Login)
		adminGroup.POST("/logout", middlewares.AdminVerifyToken(), app.Ctrl.Logout)
		adminGroup.POST("/user/list", admin.Ctrl.GetUserList)
	}
}
