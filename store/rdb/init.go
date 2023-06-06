package rdb

import (
	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/gcnf"
)

func init() {
	var rdbs = gcnf.GetConfigMap(constants.KEY_REDIS)

	InitConn(rdbs)
}
