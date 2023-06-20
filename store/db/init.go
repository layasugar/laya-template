package db

import (
	"github.com/layasugar/laya/gcnf"
)

const KEY_MYSQL = "mysql"

func init() {
	// 初始化数据库连接
	var dbs = gcnf.GetConfigMap(KEY_MYSQL)

	// 解析dbs
	InitConn(dbs)
}
