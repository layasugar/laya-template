package routers

import (
	"github.com/LaYa-op/laya-go/controllers/file"
	"github.com/LaYa-op/laya-go/controllers/hall"
	"github.com/LaYa-op/laya-go/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	// 登录注册
	r.POST("/hall/user/rgt", hall.Ctrl.Register)
	r.POST("/hall/captcha", hall.Ctrl.Captcha)
	r.POST("/hall/user/login", hall.Ctrl.Login)
	r.POST("/hall/user/tLogin", hall.Ctrl.TokenLogin)
	r.POST("/hall/send/phone", hall.Ctrl.Phone)
	r.POST("/hall/user/pwd", hall.Ctrl.EditUserPwd)

	// 文件服务器
	r.POST("/hall/files/upload", file.Ctrl.Upload)
	r.StaticFS("/hall/files", http.Dir("files"))

	authorized := r.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.POST("/hall/user/getUserInfo", hall.Ctrl.GetUserInfo)
	}
}
