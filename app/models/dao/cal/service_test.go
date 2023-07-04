// 链路追踪的测试文件

package cal_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/layasugar/laya-template/pkg/core/metautils"
	"github.com/layasugar/laya-template/routes/pb"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var RespSuc = []byte(`{
    "data": {"code": "trace-http-test"},
    "message": "操作成功",
    "status_code": 200
}`)

var agentHost = "127.0.0.1:6831"

func TestStartHttp(t *testing.T) {
	tracer, closer, _ := NewJaeger("http_server")
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
		_, _ = w.Write(RespSuc)
	})

	_ = http.ListenAndServe("0.0.0.0:10081", serverMux)
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
	pb.RegisterGreeterServer(s, &sayHello{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}

type sayHello struct {
	*pb.UnimplementedGreeterServer
}

func (ctrl *sayHello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	tracer, closer, _ := NewJaeger("grpc_server")
	defer closer.Close()

	md := metautils.ExtractIncoming(ctx)
	spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(md))
	if err != nil {
		log.Println(err.Error())
	}
	firstSpan := tracer.StartSpan("sayHello", ext.RPCServerOption(spanCtx))
	ext.HTTPUrl.Set(firstSpan, "sayHello")
	ext.HTTPMethod.Set(firstSpan, "grpc")
	firstSpan.SetTag("is_debug", "1")

	time.Sleep(50 * time.Millisecond)
	firstSpan.Finish()

	return &pb.HelloReply{Message: in.Name}, nil
}
func (ctrl *sayHello) GrpcTraceTest(ctx context.Context, in *pb.GrpcTraceTestReq) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: fmt.Sprintf("%d", in.Kind)}, nil
}

func NewJaeger(serverName string) (opentracing.Tracer, io.Closer, error) {
	var cfg = jaegerCfg.Configuration{
		ServiceName: serverName,

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
	return cfg.NewTracer(
		jaegerCfg.Logger(jLogger),
	)
}
