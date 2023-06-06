package edb

import (
	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/gcnf"
)

func init() {
	var edbs = gcnf.GetConfigMap(constants.KEY_ES)

	InitConn(edbs)
}
