package controllers

import (
	"github.com/layatips/laya/gresp"
)

// BaseController the controller with some useful and common function
var Ctrl = &BaseCtrl{}

type BaseCtrl struct {
	gresp.Resp
}
