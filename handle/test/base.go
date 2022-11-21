package test

import (
	"github.com/layasugar/laya-template/controllers"
	"github.com/layasugar/laya-template/pb"
)

var Ctrl = &controller{}

type controller struct {
	controllers.BaseCtrl
	*pb.UnimplementedGreeterServer
}
