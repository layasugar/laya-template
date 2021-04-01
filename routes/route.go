package routes

import (
	"github.com/layatips/laya"
	"github.com/layatips/laya-go/controllers/admin"
	"github.com/layatips/laya-go/controllers/app"
	"github.com/layatips/laya-go/controllers/file"
	"net/http"
)

func RegisterRoute(r *laya.WebServer) {
	// 获取用户信息
	r.POST("/app/user/info", app.Ctrl.GetUserInfo)

	// 获取用户列表
	r.POST("/admin/user/list", admin.Ctrl.GetUserList)

	// 文件服务器
	r.POST("/app/files/upload", file.Ctrl.Upload)
	r.StaticFS("/app/files", http.Dir("files"))
}
