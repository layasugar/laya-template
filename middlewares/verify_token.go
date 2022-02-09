package middlewares

import (
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya-template/global/errno"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya/genv"
)

// UserVerifyToken Middlewares
func UserVerifyToken() laya.WebHandlerFunc {
	return func(ctx *laya.WebContext) {
		// 从header头里获取 auth  然后去redis里面获取数据对比
		tokenRedisKey := genv.AppName() + global.TokenRedisKey
		key := fmt.Sprintf(tokenRedisKey, ctx.GetHeader(global.UserAuth))
		userData, err := dao.Rdb.Get(ctx, key).Result()
		if err != nil {
			ctx.WarnF("userVerifyToken rdb.Get fail,err:%s", err.Error())
			ctx.AbortWithStatusJSON(200, map[string]interface{}{
				"status_code": errno.ComUnauthorized,
				"message":     "auth verify fail",
				"request_id":  ctx.GetLogId(),
				"data":        "",
			})
		}
		ctx.Set(global.UserInfo, userData)
		ctx.Next()
	}
}
