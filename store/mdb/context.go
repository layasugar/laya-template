package mdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/layasugar/laya/store/cm"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/event"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	tSpanName = "mongo"
)

type tracer struct {
	spans sync.Map
}

func NewTracer() *tracer {
	return &tracer{}
}

const prefix = "mongodb."

func (t *tracer) HandleStartedEvent(ctx context.Context, evt *event.CommandStartedEvent) {
	if evt == nil {
		return
	}
	_, span := cm.ParseSpanByCtx(ctx, tSpanName)

	if nil != span {
		span.SetAttributes(attribute.String("db.type", "mongo"))
		span.SetAttributes(attribute.String("db.instance", evt.DatabaseName))
		span.SetAttributes(attribute.String("db.statement", evt.Command.String()))
		span.SetAttributes(attribute.String("db.host", evt.ConnectionID))
		span.SetAttributes(attribute.String("span.kind", fmt.Sprintf("%d", trace.SpanKindClient)))
		span.SetAttributes(attribute.String("component", "golang-mongo"))
	}
	t.spans.Store(evt.RequestID, span)
}

func (t *tracer) HandleSucceededEvent(ctx context.Context, evt *event.CommandSucceededEvent) {
	if evt == nil {
		return
	}
	if rawSpan, ok := t.spans.Load(evt.RequestID); ok {
		defer t.spans.Delete(evt.RequestID)
		if span, ok := rawSpan.(opentracing.Span); ok {
			defer span.Finish()
			//span.SetTag(prefix+"reply", string(evt.Reply))
			span.SetTag(prefix+"duration", evt.DurationNanos)
		}
	}
}

func (t *tracer) HandleFailedEvent(ctx context.Context, evt *event.CommandFailedEvent) {
	if evt == nil {
		return
	}
	if rawSpan, ok := t.spans.Load(evt.RequestID); ok {
		defer t.spans.Delete(evt.RequestID)
		if span, ok := rawSpan.(opentracing.Span); ok {
			defer span.Finish()
			ext.Error.Set(span, true)
			span.SetTag(prefix+"duration", evt.DurationNanos)
			span.LogFields(log.String("failure", evt.Failure))
		}
	}
}
