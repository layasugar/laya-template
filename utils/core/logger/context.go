package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Logger 日志
type Logger interface {
	LogId() string
	Debug(template string, args ...interface{})
	Info(template string, args ...interface{})
	Warn(template string, args ...interface{})
	Error(template string, args ...interface{})
	Field(key string, value interface{}) logrus.Fields
}

func (ctx *Context) Debug(template string, args ...interface{}) {
	Debug(ctx.logID, template, args...)
}

func (ctx *Context) Info(template string, args ...interface{}) {
	Info(ctx.logID, template, args...)
}

func (ctx *Context) Warn(template string, args ...interface{}) {
	Warn(ctx.logID, template, args...)
}

// ErrorF 打印程序错误日志
func (ctx *Context) Error(template string, args ...interface{}) {
	Error(ctx.logID, template, args...)
}

func (ctx *Context) Field(key string, value interface{}) logrus.Fields {
	v := fmt.Sprintf("%v", value)
	return map[string]interface{}{key: v}
}

// Context logger
type Context struct {
	logID string
}

var _ Logger = &Context{}

// NewContext new obj
func NewContext(logID string) Logger {
	ctx := &Context{
		logID: logID,
	}
	return ctx
}

// LogId 得到LogId
func (ctx *Context) LogId() string {
	return ctx.logID
}
