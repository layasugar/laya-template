package middlewares

import (
	"fmt"
	"github.com/layasugar/laya-template/app/models/dao"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/global/errno"
	"github.com/layasugar/laya-template/pkg/core"
)

// UserVerifyToken Middlewares
func UserVerifyToken() core.WebHandlerFunc {
	return func(ctx *core.Context) {
		// 从header头里获取 auth  然后去redis里面获取数据对比
		tokenRedisKey := gcnf.AppName() + global.TokenRedisKey
		key := fmt.Sprintf(tokenRedisKey, ctx.Gin().GetHeader(global.UserAuth))
		userData, err := dao.Rdb().Get(ctx, key).Result()
		if err != nil {
			ctx.Warn("userVerifyToken rdb.Get fail,err:%s", err.Error())
			ctx.Gin().AbortWithStatusJSON(200, map[string]interface{}{
				"status_code": errno.ComUnauthorized,
				"message":     "auth verify fail",
				"request_id":  ctx.LogId(),
				"data":        "",
			})
		}
		ctx.Set(global.UserInfo, userData)
		ctx.Gin().Next()
	}
}
