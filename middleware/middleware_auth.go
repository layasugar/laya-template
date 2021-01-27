package middleware

import (
	"github.com/gin-gonic/gin"
)

// implements the controllers.HandlerWrapper
func (*Middleware) Auth(c *gin.Context) {
	//token := c.GetHeader("Token")
	//uid, err := redis.Rdb.Get(context.Background(), "user:token:"+token).Result()
	//if err != nil {
	//	c.Set("$.TokenErr.code", response.TokenErr)
	//	c.Abort()
	//	return
	//}
	//
	//ID, _ := strconv.ParseInt(uid, 10, 64)
	//c.Set("uid", ID)
	//c.Next()
}
