package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/controllers/test"
	"github.com/layasugar/laya-template/pb"
)

// RegisterHttpTest 注册一组http路由
func RegisterHttpTest(r *laya.WebServer) {
	r.POST("/trace-http-test", test.Ctrl.HttpTraceTest)
}

// RegisterRpcRoutes 注册一组rpc路由
func RegisterRpcRoutes(s *laya.GrpcServer) {
	pb.RegisterGreeterServer(s.Server, test.Ctrl)
}
