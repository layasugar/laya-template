package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/controllers/test"
)

func RegisterTest(r *laya.WebServer) {
	r.POST("/trace-http-test", test.Ctrl.HttpTraceTest)
}

// RegisterRpcRoutes 注册一组rpc路由
func RegisterRpcRoutes(s *laya.PbRPCServer) {
	s.AddHandler(test.Ctrl.RpcTraceTest)
}