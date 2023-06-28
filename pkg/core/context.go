package core

import (
	"time"

	"github.com/gin-gonic/gin"
	a "github.com/layasugar/laya/core/alarm"
	"github.com/layasugar/laya/core/constants"
	d "github.com/layasugar/laya/core/data"
	l "github.com/layasugar/laya/core/logger"
	"github.com/layasugar/laya/core/metautils"
	t "github.com/layasugar/laya/core/trace"
)

// Context is the carrier of request and response
type Context struct {
	gin *gin.Context
	d.Storage
	l.Logger
	t.Tracer
	a.Alarm
}

// Deadline returns the time when work done on behalf of this contextx
// should be canceled. Deadline returns ok==false when no deadline is
// set. Successive calls to Deadline return the same results.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done returns a channel that's closed when work done on behalf of this
// contextx should be canceled. Done may return nil if this contextx can
// never be canceled. Successive calls to Done return the same value.
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err returns a non-nil error value after Done is closed,
// successive calls to Err return the same error.
// If Done is not yet closed, Err returns nil.
// If Done is closed, Err returns a non-nil error explaining why:
// Canceled if the contextx was canceled
// or DeadlineExceeded if the contextx's deadline passed.
func (c *Context) Err() error {
	return nil
}

// Value returns the value associated with this contextx for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val, _ := c.Get(keyAsString)
		return val
	}
	return nil
}

func (c *Context) Gin() *gin.Context {
	return c.gin
}

const ginFlag = "__gin__gin"

// NewContext 初始化上下文
// name uri或者spanName
// md header参数
func NewContext(st constants.SERVERTYPE, name string, md metautils.NiceMD, gin *gin.Context) *Context {
	if st == constants.SERVERGIN {
		obj, existed := gin.Get(ginFlag)
		if existed {
			return obj.(*Context)
		}
	}
	var ctx = &Context{
		Storage: d.NewContext(),
		Alarm:   a.NewContext(),
		Tracer:  t.NewTraceContext(name, md),
	}

	md.Set(constants.X_REQUESTID, ctx.TraceId())
	ctx.Set(constants.X_REQUESTID, ctx.TraceId())
	ctx.Logger = l.NewContext(ctx.TraceId())
	if st == constants.SERVERGIN {
		ctx.gin = gin
		gin.Set(ginFlag, ctx)
	}

	return ctx
}
