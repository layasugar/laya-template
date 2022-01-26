package app

import (
	"github.com/layasugar/laya-template/controllers"
)

var Ctrl = &controller{}

type controller struct {
	*controllers.BaseCtrl
}
