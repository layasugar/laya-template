package cm

import (
	"context"

	"github.com/layasugar/laya"
	"go.opentelemetry.io/otel/trace"
)

// ParseSpanByCtx 公共方法, 从ctx中获取
func ParseSpanByCtx(ctx context.Context, spanName string) (context.Context, trace.Span) {
	layaCtx, ok := ctx.(*laya.Context)
	if ok {
		return layaCtx.Start(ctx, spanName)
	}

	return nil, nil
}

// ParseLogIdByCtx 从context中解析出logId
func ParseLogIdByCtx(ctx context.Context) string {
	if webCtx, okInterface := ctx.(*laya.Context); okInterface {
		return webCtx.LogId()
	}
	return ""
}
