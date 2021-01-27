package controllers

import (
	"fmt"
	"github.com/LaYa-op/laya/glogs"
	"github.com/LaYa-op/laya/response"
	"github.com/LaYa-op/laya/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Base the controller with some useful and common function
type Base struct{}

// Success it's ok, suc response
func (bc *Base) Suc(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, bc.rspData(c, nil, data))
}

// Fail response the error info
func (bc *Base) Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, bc.rspData(c, err, nil))
}

// RawJSONString json 数据返回
func (bc *Base) RawJSONString(c *gin.Context, data string) {
	w := c.Writer
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(data))
}

// RawString raw 数据返回
func (bc *Base) RawString(c *gin.Context, data string) {
	w := c.Writer
	w.WriteHeader(200)
	_, _ = w.Write([]byte(data))
}

func (bc *Base) rspData(c *gin.Context, err error, data interface{}) interface{} {
	code, msg := 0, "success"
	if err != nil {
		code, msg = response.Exchange(err)
	}

	rsp := map[string]interface{}{
		"errno":      code,
		"errmsg":     msg,
		"request_id": utils.NewLogID(),
		"data":       data,
	}

	if v := fmt.Sprint(rsp["data"]); v == "[]" || v == "map[]" || v == "<nil>" {
		glogs.InfoF("watch_data", "%v", rsp["data"])
		rsp["data"] = struct{}{}
	}

	return rsp
}
