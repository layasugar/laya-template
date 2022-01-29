package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/controllers/test"
)

func RegisterTest(r *laya.WebServer) {
	r.POST("/trace-test", test.Ctrl.TraceTest)
}
