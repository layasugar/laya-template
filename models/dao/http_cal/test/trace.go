package test

import (
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya/gcal"
	"net/http"
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
)

var path = "/server-b/fast"
var serviceName = "trace_test"

// GetTraceTest 获取包信息
func GetTraceTest(ctx *laya.WebContext) (*Data, error) {
	req := gcal.HTTPRequest{
		Method: "POST",
		Path:   path,
		Body: map[string]string{
			"data": "success",
		},
		Ctx:       ctx,
		Header: map[string][]string{
			"Host": []string{"12312"},
		},
	}
	response := CalResp{}
	err := gcal.Cal(serviceName, req, &response, gcal.JSONConverter)

	// 状态码非 200
	if response.Head.StatusCode != http.StatusOK {
		return &response.Body.Data, errors.New("NETWORK_ERROR")
	}

	return &response.Body.Data, err
}
