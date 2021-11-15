package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya/genv"
)

// Ctrl the controllers with some useful and common function
var Ctrl = &BaseCtrl{}

type BaseCtrl struct {
	laya.Resp
}

type Memories struct {
	M Condition `json:"M"`
}

type Condition struct {
	Count int                      `json:"count"`
	Item  []map[string]interface{} `json:"item"`
}

// Version version
func (ctrl *BaseCtrl) Version(c *gin.Context) {
	res := genv.AppName() + " api version: 1.0.0\r\n" + "app_url: " + genv.AppUrl()
	_, _ = c.Writer.Write([]byte(res))
	return
}
