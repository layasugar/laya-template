// 请求测试文件

package http_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao/cal/http_test/pb_test"
	"github.com/layasugar/laya/gcal"
	"google.golang.org/grpc"
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

	RpcData struct {
		Message string `json:"message"`
	}
)

var path = "/server-b/fast"
var serviceName1 = "http_to_http_test"
var serviceName2 = "http_to_grpc_test"

// HttpToHttpTraceTest Http测试, body是interface可以发送任何类型的数据
func HttpToHttpTraceTest(ctx *laya.WebContext) (*Data, error) {
	ctx.InfoF("开始请求了, %s","aaaa")
	req := gcal.HTTPRequest{
		Method: "POST",
		Path:   path,
		Body: map[string]string{
			"data": "success",
		},
		Ctx: ctx,
		Header: map[string][]string{
			"Host": []string{"12312"},
		},
	}
	response := CalResp{}
	err := gcal.Do(serviceName1, req, &response, gcal.JSONConverter)

	// 状态码非 200
	if response.Head.StatusCode != http.StatusOK {
		return &response.Body.Data, errors.New("NETWORK_ERROR")
	}
	ctx.InfoF("结束请求了, %s","bbbb")
	return &response.Body.Data, err
}

// HttpToGrpcTraceTest grpc测试
func HttpToGrpcTraceTest(ctx *laya.WebContext) (*RpcData, error) {
	// 连接服务器
	conn, err := grpc.Dial(":10082")
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb_test.NewGreeterClient(conn)
	// 调用服务端的SayHello
	r, err := c.SayHello(context.Background(), &pb_test.HelloRequest{Name: "q1mi"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s !\n", r.Message)
	return nil, err
}
