package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/LaYa-op/laya"
	r "github.com/LaYa-op/laya/response"
	"strconv"
)

// implements the controllers.HandlerWrapper
func Auth(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, err := ship.Redis.Get("user:token:" + token).Result()
	if err != nil {
		c.Set("$.TokenErr.code", r.TokenErr)
		c.Abort()
		return
	}

	ID, _ := strconv.ParseInt(uid, 10, 64)
	c.Set("uid", ID)
	c.Next()
}
