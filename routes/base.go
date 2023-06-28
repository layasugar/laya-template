package routes

import (
	"net/http"

	"github.com/layasugar/laya-template/app/handler/file"
	"github.com/layasugar/laya-template/pkg/core"
)

func Register(r *core.WebServer) {
	r.GET("/", file.Ctrl.Version)   // version
	r.POST("/test", file.Ctrl.Test) // 测试接口

	// 文件服务器
	r.POST("/app/files/upload", file.Ctrl.Upload)
	r.StaticFS("/app/files", http.Dir("files"))

	RegisterApp(r)
	RegisterAdmin(r)
	RegisterHttpTest(r)
}
