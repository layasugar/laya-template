package routes

import (
	"net/http"

	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/handle/file"
)

func Register(r *laya.WebServer) {
	r.GET("/", handler.Ctrl.Version)   // version
	r.POST("/test", handler.Ctrl.Test) // 测试接口

	// 文件服务器
	r.POST("/app/files/upload", file.Ctrl.Upload)
	r.StaticFS("/app/files", http.Dir("files"))

	RegisterApp(r)
	RegisterAdmin(r)
	RegisterHttpTest(r)
}
