package db

import (
	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/gcnf"
)

func init() {
	// 初始化数据库连接和redis连接
	var dbs = gcnf.GetConfigMap(constants.KEY_MYSQL)

	// 解析dbs
	InitConn(dbs)
}
