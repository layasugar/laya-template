package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/controllers/admin"
	"github.com/layasugar/laya-template/controllers/app"
	"github.com/layasugar/laya-template/middlewares"
)

func RegisterAdmin(r *laya.WebServer) {
	adminGroup := r.Group("/admin/v1")
	{
		adminGroup.POST("/login", admin.Ctrl.Login)
		adminGroup.POST("/logout", middlewares.AdminVerifyToken(), app.Ctrl.Logout)
		adminGroup.POST("/user/list", admin.Ctrl.GetUserList)
	}
}
