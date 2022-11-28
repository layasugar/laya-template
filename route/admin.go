package route

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/middleware"

	"github.com/layasugar/laya-template/handle/admin"
	"github.com/layasugar/laya-template/handle/app"
)

func RegisterAdmin(r *laya.WebServer) {
	adminGroup := r.Group("/admin/v1")
	{
		adminGroup.POST("/login", admin.Ctrl.Login)
		adminGroup.POST("/logout", middleware.AdminVerifyToken(), app.Ctrl.Logout)
		adminGroup.POST("/user/list", admin.Ctrl.GetUserList)
	}
}
