package test

import (
	"github.com/layasugar/laya-template/handle"
	"github.com/layasugar/laya-template/pb"
)

var Ctrl = &controller{}

type controller struct {
	*handle.BaseCtrl
	*pb.UnimplementedGreeterServer
}
