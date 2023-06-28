package trace

import (
	"github.com/layasugar/laya/core/logger"
	"github.com/layasugar/laya/core/trace/b3propagator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func newZkTracer(service, url, appMod string, mod float64) {
	exporter, err := zipkin.New(
		url,
	)
	if err != nil {
		logger.Error("none", "zipkin初始化失败，err: %v", err)
		return
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(mod)),
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", appMod),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(b3propagator.New())
}
