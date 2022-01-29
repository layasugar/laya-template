package rpc_cal

import (
	"errors"
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya/gcal"
	"net/http"
	"net/url"
)

type JsonStruct struct {
	Test string `json:"test"`
}

var (
	_GetTestDataPath = "test"
	// StubTestCal function for testing
	StubTestCal func(service string, req gcal.HTTPRequest, resp *CalResp, ct gcal.ConverterType) error
)

// CalResp response type for TestCal
type CalResp struct {
	Head gcal.HTTPHead
	Body JsonStruct
}

// GetPkgData 获取包信息
func GetPkgData(ctx *laya.WebContext, query url.Values, bodyData []byte) (*JsonStruct, error) {
	path := fmt.Sprintf("%s?%s", _GetTestDataPath, query.Encode())
	req := gcal.HTTPRequest{
		Method: "POST",
		Path:   path,
		Body: map[string]string{
			"data": string(bodyData),
		},
		Header: http.Header{
			"Host": []string{"https://layasugar.cn"},
		},
		Ctx: ctx,
	}
	response := CalResp{}
	var err error
	if StubTestCal == nil {
		err = gcal.Cal("aps", req, &response, gcal.JSONConverter)
	} else {
		err = StubTestCal("aps", req, &response, gcal.JSONConverter)
	}

	// 状态码非 200
	if response.Head.StatusCode != http.StatusOK {
		return &response.Body, errors.New("NETWORK_ERROR")
	}

	return &response.Body, err
}
