package controllers

import (
	"github.com/layatips/laya/gresp"
)

// BaseCtrl the controllers with some useful and common function
var Ctrl = &BaseCtrl{}

type BaseCtrl struct {
	gresp.Resp
}
