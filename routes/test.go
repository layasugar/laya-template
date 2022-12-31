package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/handle/pb"

	"github.com/layasugar/laya-template/handle/test"
)

// RegisterHttpTest 注册一组http路由
func RegisterHttpTest(r *laya.WebServer) {
	r.POST("/full-test", test.Ctrl.FullTest)
}

// RegisterRpcRoutes 注册一组rpc路由
func RegisterRpcRoutes(s *laya.GrpcServer) {
	pb.RegisterGreeterServer(s.Server, test.Ctrl)
}
