package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya/glogs"
)

func SetTrace(c *gin.Context) {
	if glogs.Tracer != nil {
		if c.Request.RequestURI != "/ready" && c.Request.RequestURI != "/healthz" {
			span := glogs.StartSpanFromReq(c.Request, glogs.Tracer, c.Request.RequestURI)
			span.Tag("request-id", c.GetHeader(glogs.RequestIDName))
			_ = glogs.InjectToReq(context.WithValue(context.Background(), glogs.GetSpanContextKey(), span.Context()), c.Request)
			c.Next()
			span.Finish()
		}
	} else {
		c.Next()
	}
}
