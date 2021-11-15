//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/glogs"
	"github.com/layasugar/laya/gconf"
	"github.com/layasugar/laya/gstore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbKey    = "mysql"
	rdbKey   = "redis"
	mdbKey   = "mongo"
	kafkaKey = "kafka"
	esKey    = "es"
)

// DB is sql *db
var DB *gorm.DB

// Rdb is redis *client
var Rdb *redis.Client

func Init() {
	//mysql
	dbCfg := &gorm.Config{Logger: glogs.Default(glogs.Sugar, logger.Info)}
	dbPoolCfg := gstore.DbPoolCfg{}
	DB = gstore.InitDB(gconf.C.GetString("mysql.dsn"), gstore.SetPoolConfig(dbPoolCfg), gstore.SetGormConfig(dbCfg))
}
