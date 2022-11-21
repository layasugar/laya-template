package middleware

import (
	"github.com/layasugar/laya"
	"log"
)

// TestMiddleware 测试web中间件
func TestMiddleware() laya.WebHandlerFunc {
	return func(ctx *laya.WebContext) {
		log.Print("测试web中间件")
	}
}
