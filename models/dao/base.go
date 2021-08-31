//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/glogs"
	"github.com/layasugar/kafka-go"
	"github.com/layasugar/laya/gcache"
	"github.com/layasugar/laya/gconf"
	"github.com/layasugar/laya/gstore"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
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

// Mdb is mongo client
var Mdb *mongo.Client

// Mem
var Mem *gcache.Cache

// kafka
var Kafka *gkafka.Engine

// es
var Es *elastic.Client

func Init() {
	// mysql init client
	// dsn 是必须配置的，其他都是非必须的
	// gstore.SetPoolConfig(dc.DbPoolCfg), gstore.SetGormConfig(dbCfg) 都是非必须的
	dc, err := gconf.GetDBConf(dbKey)
	if err != nil {
		panic(err.Error())
	}
	dbCfg := &gorm.Config{
		Logger: glogs.Default(glogs.Sugar, logger.Info),
	}
	DB = gstore.InitDB(dc.Dsn, gstore.SetPoolConfig(dc.DbPoolCfg), gstore.SetGormConfig(dbCfg))

	//redis
	rc, err := gconf.GetRdbConf(rdbKey)
	if err != nil {
		panic(err.Error())
	}
	Rdb = gstore.InitRdb(rc.DB, rc.PoolSize, rc.MaxRetries, rc.IdleTimeout, rc.Addr, rc.Pwd)

	//mongodb
	mc, err := gconf.GetMdbConf(mdbKey)
	if err != nil {
		panic(err.Error())
	}
	Mdb = gstore.InitMdb(mc.MinPoolSize, mc.MaxPoolSize, mc.DSN)
}
