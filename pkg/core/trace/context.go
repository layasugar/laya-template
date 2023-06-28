package trace

import (
	"context"

	"github.com/layasugar/laya/core/metautils"
	"github.com/layasugar/laya/core/util"
	uuid "github.com/satori/go.uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Tracer 链路
type Tracer interface {
	TraceId() string
	End(span trace.Span)
	Start(ctx context.Context, spanName string) (context.Context, trace.Span)
	Inject(c context.Context, md metautils.NiceMD)
	Extract(md metautils.NiceMD) (context.Context, trace.Span)
	TopSpan() trace.Span
}

// Context trace
type Context struct {
	traceId string
	topCtx  context.Context
	topSpan trace.Span
}

func (ctx *Context) TraceId() string {
	return ctx.traceId
}

func (ctx *Context) End(span trace.Span) {
	if nil != span {
		span.End()
	}
}

func (ctx *Context) Start(c context.Context, spanName string) (context.Context, trace.Span) {
	if t := getTracer(); t != nil {
		if ctx == nil {
			return t.Start(ctx.topCtx, spanName, trace.WithSpanKind(trace.SpanKindServer))
		} else {
			return t.Start(c, spanName, trace.WithSpanKind(trace.SpanKindServer))
		}
	}
	return context.Background(), nil
}

// Inject 将span注入到request
func (ctx *Context) Inject(c context.Context, md metautils.NiceMD) {
	if t := getTracer(); t != nil {
		if c == nil {
			otel.GetTextMapPropagator().Inject(ctx.topCtx, propagation.HeaderCarrier(md))
		} else {
			otel.GetTextMapPropagator().Inject(c, propagation.HeaderCarrier(md))
		}
	}
}

func (ctx *Context) Extract(md metautils.NiceMD) (context.Context, trace.Span) {
	c := otel.GetTextMapPropagator().Extract(context.Background(), propagation.HeaderCarrier(md))
	span := trace.SpanFromContext(c)
	return c, span
}

func (ctx *Context) TopSpan() trace.Span {
	return ctx.topSpan
}

// NewTraceContext new traceCtx
func NewTraceContext(name string, md metautils.NiceMD) *Context {
	ctx := &Context{}
	if t := getTracer(); t != nil {
		if len(md) == 0 {
			ctx.topCtx, ctx.topSpan = t.Start(context.Background(), name, trace.WithSpanKind(trace.SpanKindServer))
		} else {
			ctx.topCtx, ctx.topSpan = ctx.Extract(md)
			if ctx.topCtx == context.Background() {
				ctx.topCtx, ctx.topSpan = t.Start(context.Background(), name, trace.WithSpanKind(trace.SpanKindServer))
			}
		}
		ctx.traceId = ctx.topSpan.SpanContext().TraceID().String()
	} else {
		ctx.traceId = util.Md5(uuid.NewV4().String())
	}
	return ctx
}
