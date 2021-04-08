package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya/glogs"
	"github.com/layatips/laya/gutils"
)

func SetTrace(c *gin.Context) {
	if glogs.Tracer != nil {
		if !gutils.InSliceString(c.Request.RequestURI, gutils.IgnoreRoutes) {
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
