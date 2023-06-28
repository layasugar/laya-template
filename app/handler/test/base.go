package test

import (
	"github.com/layasugar/laya-template/app/handler"
	"github.com/layasugar/laya-template/routes/pb"
)

var Ctrl = &controller{}

type controller struct {
	*handler.BaseCtrl
	*pb.UnimplementedGreeterServer
}
