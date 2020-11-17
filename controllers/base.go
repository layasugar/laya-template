package controllers

import (
	"github.com/gin-gonic/gin"
)

// BaseController the controller with some useful and common function
type BaseController struct{}

// Suc it's ok, suc response
func (bc *BaseController) Suc(ctx *gin.Context, data interface{}) {
	ctx.Set("$.Suc.response", data)
}

// RawJSONString json 数据返回
func (bc *BaseController) RawJSONString(ctx *gin.Context, data string) {
	w := ctx.Writer
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(data))
}

// Fail response the error info
func (bc *BaseController) Fail(ctx *gin.Context, err error) {
	ctx.Set("$.Fail.code", err)
}
