package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/models/dao"
	"strconv"
)

// implements the controllers.HandlerWrapper
func (*Middleware) Test(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, err := dao.Rdb.Get(context.Background(), "user:token:"+token).Result()
	if err != nil {
		c.Set("$.TokenErr.code", errors.New("asdasd"))
		c.Abort()
		return
	}

	ID, _ := strconv.ParseInt(uid, 10, 64)
	c.Set("uid", ID)
	c.Next()
}
