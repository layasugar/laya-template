package errno

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya/genv"
	"github.com/layasugar/glogs"
	"github.com/layasugar/laya/gutils"
	"net/http"
)

const requestIDName = glogs.RequestIDName

type Resp struct{}

type response struct {
	StatusCode uint32      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	RequestID  string      `json:"request_id"`
}

func (res *Resp) Suc(c *gin.Context, data interface{}, msg ...string) {
	rr := new(response)
	rr.StatusCode = http.StatusOK
	if len(msg) == 0 {
		rr.Message = "success"
	} else {
		for _, v := range msg {
			rr.Message += "," + v
		}
	}
	rr.Data = data
	rr.RequestID = c.GetHeader(requestIDName)
	if !gutils.InSliceString(c.Request.RequestURI, gutils.IgnoreRoutes) {
		if genv.ParamLog() {
			log, _ := json.Marshal(&rr)
			glogs.InfoF(c, "出参打印", fmt.Sprintf("path=%s,content=%s", c.Request.RequestURI, log))
		}
	}
	c.JSON(http.StatusOK, &rr)
}

func (res *Resp) Fail(c *gin.Context, err error) {
	rr := new(response)
	switch err.(type) {
	case *RspError:
		rr.StatusCode, rr.Message = err.(*RspError).Render()
	default:
		rr.StatusCode = 400
		rr.Message = err.Error()
	}
	rr.RequestID = c.GetHeader(requestIDName)
	if !gutils.InSliceString(c.Request.RequestURI, gutils.IgnoreRoutes) {
		if genv.ParamLog() {
			log, _ := json.Marshal(&rr)
			glogs.InfoF(c, "出参打印", fmt.Sprintf("title=出参打印,path=%s,content=%s", c.Request.RequestURI, log))
		}
	}
	c.JSON(http.StatusOK, &rr)
}

// RawJSONString json 数据返回
func (res *Resp) RawJSONString(c *gin.Context, data string) {
	w := c.Writer
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(data))
}

// RawString raw 数据返回
func (res *Resp) RawString(c *gin.Context, data string) {
	w := c.Writer
	w.WriteHeader(200)
	_, _ = w.Write([]byte(data))
}
