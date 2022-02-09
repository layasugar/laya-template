// 请求测试文件

package rpc_test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya/gcal"
)

type (
	CalResp struct {
		Head gcal.HTTPHead
		Body Body
	}

	Body struct {
		StatusCode uint32 `json:"status_code"`
		Message    string `json:"message"`
		Data       Data   `json:"data"`
		RequestID  string `json:"request_id"`
	}

	Data struct {
		Code string `json:"code"`
	}

	RpcData struct {
		Message string `json:"message"`
	}
)

var path = "/server-b/fast"
var serviceName1 = "http_to_http_test"
var serviceName2 = "http_to_rpc_test"

// HttpTraceTest Http测试, body是interface可以发送任何类型的数据
func HttpTraceTest(ctx *laya.PbRPCContext) (*Data, error) {

	return nil, nil
}

// RpcTraceTest rpc测试
func RpcTraceTest(ctx *laya.PbRPCContext) (*RpcData, error) {

	var res = RpcData{
		Message: "1",
	}

	return &res, nil
}
