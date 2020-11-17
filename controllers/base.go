package controllers

import (
	"gitlab.xthktech.cn/bs/gxe"
	"gitlab.xthktech.cn/xthk-online/xthk_student_center_pub/library/errno"
	"strconv"
)

var tokenMap = map[string]string{
	"test": "12345678",
}

const (
	ValidTimeSapn = 30
	//token 中间件保存的用户信息键
	UserInfoKey     = "userStudentInfo"
	DefaultPage     = 1
	DefaultPageSize = 10
)

// BaseController the controller with some useful and common function
type BaseController struct{}

var zero = struct{}{}

func (bc *BaseController) rspData(ctx *gxe.WebContext, err error, data interface{}, intErrno bool, extraInfo ...map[string]interface{}) interface{} {
	code, msg := "200", "succ"
	var extra map[string]string
	if err != nil {
		code, msg, _, extra = errno.Exchange(err)
	}

	rsp := map[string]interface{}{
		"status_code": code,
		"message":     msg,
		"request_id":  ctx.GetLogID(),
		"data":        data,
	}
	if extra != nil {
		rsp["tipoption"] = extra
	}

	ctx.ErrNo = code

	errnoInt, _ := strconv.Atoi(code)
	if intErrno {
		rsp["status_code"] = errnoInt
	}

	if len(extraInfo) == 1 && extraInfo[0] != nil {
		for k, v := range extraInfo[0] {
			rsp[k] = v
		}
	}

	//if v := fmt.Sprint(rsp["data"]); v == "[]" || v == "map[]" || v == "<nil>" {
	//	ctx.AddNoticeF("watch_data", "%v", rsp["data"])
	//	rsp["data"] = zero
	//}

	return rsp
}

// Succ it's ok, succ response
func (bc *BaseController) Succ(ctx *gxe.WebContext, data interface{}, extraInfo ...map[string]interface{}) {
	ctx.JSON(200, bc.rspData(ctx, nil, data, true, extraInfo...))
}

// RawJSONString json 数据返回
func (bc *BaseController) RawJSONString(ctx *gxe.WebContext, data string) {
	w := ctx.Writer
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
}

// Fail response the error info
func (bc *BaseController) Fail(ctx *gxe.WebContext, err error, extraInfo ...map[string]interface{}) {
	ctx.AbortWithStatusJSON(200, bc.rspData(ctx, err, nil, true, extraInfo...))
}
