// 链路追踪

package trace

import (
	"github.com/layasugar/laya-template/pkg/core/constants"
	"github.com/layasugar/laya-template/pkg/gcnf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// trace 全局单例变量
var tracer trace.Tracer

// getTracer 获取全局的tracer
func getTracer() trace.Tracer {
	if nil == tracer {
		if gcnf.TraceMod() != 0 {
			switch gcnf.TraceType() {
			case constants.TRACETYPEZIPKIN:
				newZkTracer(gcnf.AppName(), gcnf.TraceAddr(), gcnf.AppMode(), gcnf.TraceMod())
			case constants.TRACETYPEJAEGER:
				newJTracer(gcnf.AppName(), gcnf.TraceAddr(), gcnf.AppMode(), gcnf.TraceMod())
			}
		}
	}

	tracer = otel.Tracer(gcnf.AppName())
	return tracer
}
