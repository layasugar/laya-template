package rdb

import (
	"github.com/layasugar/laya-template/pkg/core/constants"
	"github.com/layasugar/laya/gcnf"
)

func init() {
	var rdbs = gcnf.GetConfigMap(constants.KEY_REDIS)

	InitConn(rdbs)
}
