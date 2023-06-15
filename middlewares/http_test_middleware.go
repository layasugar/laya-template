package middlewares

import (
	"log"

	"github.com/layasugar/laya"
)

// TestMiddleware 测试web中间件
func TestMiddleware() laya.WebHandlerFunc {
	return func(ctx *laya.Context) {
		log.Print("测试web中间件")
	}
}
