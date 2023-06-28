// Package protocol 提供了 HTTP、HTTPS、NSHead、ProtoBuffer 协议支持
package protocol

import (
	"errors"
	"fmt"
	"strings"

	"github.com/layasugar/laya/gcal/context"
	"github.com/layasugar/laya/gcal/service"
)

// Protocoler 协议的接口
// 协议本身只完成数据请求
type Protocoler interface {
	Do(ctx *context.Context, addr string) (*Response, error)
	Protocol() string
}

var (
	_ Protocoler = &HTTPProtocol{}
)

// NewProtocol 创建协议
func NewProtocol(ctx *context.Context, serv service.Service, req interface{}) (p Protocoler, err error) {
	tmp, ok := req.(HTTPRequest)
	if !ok {
		return nil, fmt.Errorf("bad request type: %T", req)
	}

	protocolName := serv.GetProtocol()
	if protocolName == "" {
		as := strings.Split(tmp.CustomAddr, "://")
		if len(as) == 2 {
			protocolName = as[0]
		} else {
			return nil, errors.New("protocol is nil")
		}
	}
	if protocolName == "http" || protocolName == "https" {
		return NewHTTPProtocol(ctx, serv, &tmp, protocolName == "https")
	}

	return nil, fmt.Errorf("unknow protocol: %s ", protocolName)
}

// Response 通用的返回
type Response struct {
	Body      interface{}
	Head      interface{}
	Request   interface{}
	OriginRsp interface{}
}
