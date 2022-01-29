package test_test

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"log"
	"net/http"
	"testing"
	"time"
)

var RespSuc = []byte(`{
    "data": {"code": "trace-test"},
    "message": "操作成功",
    "status_code": 200
}`)

var agentHost = "127.0.0.1:6831"

// TestStartServiceA 入口创建第一个span
func TestStartServiceA(t *testing.T) {
	var cfg = jaegerCfg.Configuration{
		ServiceName: "service-a",
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
	serverMux.HandleFunc("/server-a/fast", func(w http.ResponseWriter, r *http.Request) {
		// operationName 操作名
		// 第一个span
		firstSpan := tracer.StartSpan("server-a")
		firstSpan.SetTag("入参", "123")
		firstSpan.SetTag("db.type", "mysql")
		firstSpan.SetTag("is_debug", "1")
		firstSpan.SetTag("出参", string(RespSuc))

		// 子span,查询sql
		childSpan := tracer.StartSpan(
			"mysql",
			opentracing.FollowsFrom(firstSpan.Context()),
		)
		childSpan.SetTag("sql1", "select xxxxxxx")
		childSpan.SetTag("is_debug", "1")
		time.Sleep(100 * time.Millisecond)
		childSpan.Finish()

		// 发起请求
		url := "http://127.0.0.1:10082/server-b/fast"
		req, _ := http.NewRequest("GET", url, nil)
		tracer.Inject(firstSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		ext.HTTPUrl.Set(firstSpan, url)
		ext.HTTPMethod.Set(firstSpan, "GET")
		resp, _ := http.DefaultClient.Do(req)
		_ = resp

		w.Write(RespSuc)
		firstSpan.Finish()
	})

	http.ListenAndServe("0.0.0.0:10081", serverMux)
}

func TestStartServiceB(t *testing.T) {
	var cfg = jaegerCfg.Configuration{
		ServiceName: "service-b",

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
		firstSpan.SetTag("is_debug", "1")
		// 发起请求
		//url := "http://127.0.0.1:10083/server-c/fast"
		//req, _ := http.NewRequest("GET", url, nil)
		////req.Header.Set("x-b3-traceid", traceID)
		//tracer.Inject(firstSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		//ext.HTTPUrl.Set(firstSpan, url)
		//ext.HTTPMethod.Set(firstSpan, "GET")

		//resp, _ := http.DefaultClient.Do(req)
		//_ = resp
		firstSpan.Finish()
		w.Write(RespSuc)
	})

	http.ListenAndServe("0.0.0.0:10082", serverMux)
}

func TestStartServiceC(t *testing.T) {
	var cfg = jaegerCfg.Configuration{
		ServiceName: "service-c",
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
	serverMux.HandleFunc("/server-c/fast", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Println(err.Error())
		}
		firstSpan := tracer.StartSpan(r.URL.Path, ext.RPCServerOption(spanCtx))
		time.Sleep(100 * time.Millisecond)

		// 子span,查询sql
		childSpan := tracer.StartSpan(
			"mysql",
			opentracing.FollowsFrom(firstSpan.Context()),
		)
		childSpan.SetTag("sql1", "select xxxxxxx")
		time.Sleep(100 * time.Millisecond)
		childSpan.Finish()

		// 发起请求
		url := "http://127.0.0.1:10084/server-d/fast"
		req, _ := http.NewRequest("GET", url, nil)
		tracer.Inject(firstSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		ext.HTTPUrl.Set(firstSpan, url)
		ext.HTTPMethod.Set(firstSpan, "GET")

		resp, _ := http.DefaultClient.Do(req)
		_ = resp
		w.Write(RespSuc)

		w.Write(RespSuc)
		firstSpan.Finish()
	})

	http.ListenAndServe("0.0.0.0:10083", serverMux)
}

func TestStartServiceD(t *testing.T) {
	var cfg = jaegerCfg.Configuration{
		ServiceName: "service-d",
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
	serverMux.HandleFunc("/server-d/fast", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Println(err.Error())
		}
		startSpan := tracer.StartSpan(r.URL.Path, ext.RPCServerOption(spanCtx))

		w.Write(RespSuc)
		startSpan.Finish()
	})

	http.ListenAndServe("0.0.0.0:10084", serverMux)
}
