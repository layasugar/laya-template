package middlewares

import (
	"log"

	"github.com/layasugar/laya-template/pkg/core"
)

// TestMiddleware 测试web中间件
func TestMiddleware() core.WebHandlerFunc {
	return func(ctx *core.Context) {
		log.Print("测试web中间件")
	}
}
