package router

import (
	"github.com/gin-gonic/gin"
	"laya-go/server/hall/handler"
	"laya-go/ship/middleware"
	"net/http"
)

func Init(r *gin.Engine) {
	// 登录注册
	r.POST("/hall/user/rgt", handler.Register)
	r.POST("/hall/captcha", handler.Captcha)
	r.POST("/hall/user/login", handler.Login)
	r.POST("/hall/user/tLogin", handler.TokenLogin)
	r.POST("/hall/send/phone", handler.Phone)
	r.POST("/hall/user/pwd", handler.EditUserPwd)

	// 文件服务器
	r.POST("/hall/files/upload", handler.Upload)
	r.StaticFS("/hall/files", http.Dir("files"))

	authorized := r.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.POST("/hall/user/getUserInfo", handler.GetUserInfo)
	}
}
