package middleware

import (
	"github.com/LaYa-op/laya"
	"github.com/LaYa-op/laya/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// implements the controllers.HandlerWrapper
func Auth(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, err := laya.Redis.Get("user:token:" + token).Result()
	if err != nil {
		c.Set("$.TokenErr.code", response.TokenErr)
		c.Abort()
		return
	}

	ID, _ := strconv.ParseInt(uid, 10, 64)
	c.Set("uid", ID)
	c.Next()
}
