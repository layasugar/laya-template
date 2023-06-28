package trace

import (
	"github.com/layasugar/laya/core/logger"
	"github.com/layasugar/laya/core/trace/jaegerpropagetor"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func newJTracer(service, addr, appMod string, mod float64) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(addr)))
	if err != nil {
		logger.Error("app", "jaeger初始化失败，err: %v", err)
		return
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(mod)),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", appMod),
		)),
	)
	otel.SetTracerProvider(tp)
	p := jaegerpropagetor.Jaeger{}
	otel.SetTextMapPropagator(p)
	logger.Debug("app", "tracer success")
}
