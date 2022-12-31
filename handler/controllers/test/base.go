package test

import (
	"github.com/layasugar/laya-template/handle/pb"
)

var Ctrl = &controller{}

type controller struct {
	*handler.BaseCtrl
	*pb.UnimplementedGreeterServer
}
