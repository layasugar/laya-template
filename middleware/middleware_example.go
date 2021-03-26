package middleware

import (
	"context"
	"github.com/layatips/laya-go/models/dao/rdb"
	"github.com/layatips/laya/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// implements the controllers.HandlerWrapper
func (*Middleware) Test(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, err := rdb.Dao.Get(context.Background(), "user:token:"+token).Result()
	if err != nil {
		c.Set("$.TokenErr.code", response.TokenErr)
		c.Abort()
		return
	}

	ID, _ := strconv.ParseInt(uid, 10, 64)
	c.Set("uid", ID)
	c.Next()
}
