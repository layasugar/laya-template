package global

import (
	"net/http"

	"github.com/layasugar/laya/gcnf"
)

type HttpResp struct{}

type Response struct {
	Code       uint32      `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	XRequestID string      `json:"x_request_id"`
}

// rspError 错误处理
type rspError struct {
	Code uint32
	Msg  string
}

func (re *rspError) Error() string {
	return re.Msg
}

func Err(code uint32) (err error) {
	err = &rspError{
		Code: code,
	}
	return err
}

// Render 渲染
func (re *rspError) render() (uint32, string) {
	msg := gcnf.LoadErrMsg(re.Code)
	if msg == "" {
		msg = "sorry, system err"
	}
	re.Msg = msg
	return re.Code, re.Msg
}

func (res *HttpResp) Suc(ctx *core.Context, data interface{}, msg ...string) {
	rr := new(Response)
	rr.Code = http.StatusOK
	if len(msg) == 0 {
		rr.Msg = "success"
	} else {
		for _, v := range msg {
			if rr.Msg != "" {
				rr.Msg += ", " + v
			} else {
				rr.Msg += v
			}
		}
	}
	rr.Data = data
	rr.XRequestID = ctx.LogId()
	ctx.Gin().JSON(http.StatusOK, &rr)
}

func (res *HttpResp) Fail(ctx *core.Context, err error) {
	rr := new(Response)
	switch t := err.(type) {
	case *rspError:
		rr.Code, rr.Msg = t.render()
	default:
		rr.Code = 400
		rr.Msg = err.Error()
	}
	rr.XRequestID = ctx.LogId()
	ctx.Gin().JSON(http.StatusOK, &rr)
}

// RawJSONString json 数据返回
func (res *HttpResp) RawJSONString(ctx *core.Context, data string) {
	w := ctx.Gin().Writer
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(data))
}

// RawString raw 数据返回
func (res *HttpResp) RawString(ctx *core.Context, data string) {
	w := ctx.Gin().Writer
	w.WriteHeader(200)
	_, _ = w.Write([]byte(data))
}
