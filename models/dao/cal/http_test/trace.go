// 请求测试文件

package http_test

import (
	"errors"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/pb"
	"github.com/layasugar/laya/gcal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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
var serviceName1 = "http_test"
var serviceName2 = "grpc_test"

// HttpToHttpTraceTest Http测试, body是interface可以发送任何类型的数据
func HttpToHttpTraceTest(ctx *laya.WebContext) (*Data, error) {
	log.Printf("开始请求了, %s", "aaaa")
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
		Converter: gcal.JSONConverter,
	}
	response := CalResp{}
	err := gcal.Do(serviceName1, req, &response, gcal.JSONConverter)

	// 状态码非 200
	if response.Head.StatusCode != http.StatusOK {
		return &response.Body.Data, errors.New("NETWORK_ERROR")
	}
	log.Printf("结束请求了, %s", "bbbb")
	return &response.Body.Data, err
}

// HttpToGrpcTraceTest grpc测试
func HttpToGrpcTraceTest(ctx *laya.WebContext) (*RpcData, error) {
	conn, errDial := grpc.Dial("127.0.0.1:9601", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errDial != nil {
		return nil, errDial
	}

	var sex int32 = 100
	d := pb.NewUserClient(conn)
	res, err := d.SayHello(ctx, &pb.HiUser{Name: "xxxx", Sex: &sex})
	if err != nil {
		return nil, err
	}
	return &RpcData{
		Message: res.Message,
	}, err

	//conn := gcal.GetRpcConn(serviceName2)
	//if conn == nil {
	//	return nil, errors.New("连接不存在")
	//}
	//
	//c := pb.NewGreeterClient(conn)
	//
	//res, err := c.SayHello(ctx, &pb.HelloRequest{Name: "q1mi"})
	//if err != nil {
	//	return nil, err
	//}
	//return &RpcData{
	//	Message: res.Message,
	//}, err
}
