// 链路追踪的测试文件

package http_test_test

import (
	"context"
	"fmt"
	"github.com/layasugar/laya-template/models/dao/cal/http_test/pb_test"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"testing"
)

var RespSuc = []byte(`{
    "data": {"code": "trace-http-test"},
    "message": "操作成功",
    "status_code": 200
}`)

var agentHost = "127.0.0.1:6831"
var Tracer opentracing.Tracer

func TestStartHttp(t *testing.T) {
	var cfg = jaegerCfg.Configuration{
		ServiceName: "service-http",

		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: agentHost,
		},
	}
	jLogger := jaegerLog.StdLogger
	tracer, closer, _ := cfg.NewTracer(
		jaegerCfg.Logger(jLogger),
	)

	defer closer.Close()

	var serverMux = http.NewServeMux()
	serverMux.HandleFunc("/server-b/fast", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Println(err.Error())
		}
		firstSpan := tracer.StartSpan(r.URL.Path, ext.RPCServerOption(spanCtx))
		ext.HTTPUrl.Set(firstSpan, r.URL.Path)
		ext.HTTPMethod.Set(firstSpan, r.Method)
		firstSpan.SetTag("is_debug", "1")
		firstSpan.Finish()
		w.Write(RespSuc)
	})

	http.ListenAndServe("0.0.0.0:10081", serverMux)
}

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb_test.HelloRequest) (*pb_test.HelloReply, error) {
	var a = in.Name
	return &pb_test.HelloReply{Message: "hello " + a}, nil
}

func TestStartRpcx(t *testing.T) {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":10082")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	pb_test.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
