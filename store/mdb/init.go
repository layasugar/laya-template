package mdb

import (
	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/gcnf"
)

func init() {
	var mdbs = gcnf.GetConfigMap(constants.KEY_MONGO)
	InitConn(mdbs)
}
