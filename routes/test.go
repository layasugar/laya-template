package routes

import (
	"github.com/layasugar/laya-template/app/handler/test"
	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/routes/pb"
)

// RegisterHttpTest 注册一组http路由
func RegisterHttpTest(r *core.WebServer) {
	r.POST("/full-test", test.Ctrl.FullTest)
}

// RegisterRpcRoutes 注册一组rpc路由
func RegisterRpcRoutes(s *core.GrpcServer) {
	pb.RegisterGreeterServer(s.Server, test.Ctrl)
}
