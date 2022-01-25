//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/laya/gconf"
	"github.com/layasugar/laya/gstore"
	"gorm.io/gorm"
)

// DB is sql *db
var DB *gorm.DB

// Rdb is redis *client
var Rdb *redis.Client

func Init() {
	// mysql
	DB = gstore.InitDB(gconf.V.GetString("mysql.dsn"), gstore.LevelInfo)

	// redis
}
