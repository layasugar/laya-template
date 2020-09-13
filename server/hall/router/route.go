package router

import (
	"github.com/gin-gonic/gin"
	"laya-go/server/hall/handler"
	"laya-go/ship/middleware"
)

func Init(r *gin.Engine) {
	// 登录注册
	//r.POST("/hall/user/rgt", handler.Register)
	//r.POST("/hall/captcha", handler.Captcha)
	r.POST("/hall/user/login", handler.Login)
	//r.POST("/hall/user/tLogin", handler.TokenLogin)
	//r.POST("/hall/send/phone", handler.Phone)
	//r.POST("/hall/user/pwd", handler.EditUserPwd)
	//
	//// 文件服务器
	//r.POST("/hall/files/upload", handler.Upload)
	//r.StaticFS("/hall/files", http.Dir("files"))

	// 支付回调
	//r.POST("/hall/cash/ttNotify", pay.TTNotify)
	//r.POST("/hall/notify/:channel", handler.Notify)

	authorized := r.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.POST("/hall/user/getUserInfo", handler.GetUserInfo)
		//authorized.POST("/hall/user/payOrder", handler.CreateOrder)
	}
}
