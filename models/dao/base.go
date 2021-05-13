//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layatips/laya/gcache"
	"github.com/layatips/laya/gconf"
	"github.com/layatips/laya/gkafka"
	"github.com/layatips/laya/glogs"
	"github.com/layatips/laya/gstore"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbKey    = "mysql"
	rdbKey   = "redis"
	mdbKey   = "mongo"
	kafkaKey = "kafka"
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

func Init() {
	//mysql
	dc, err := gconf.GetDBConf(dbKey)
	if err != nil {
		panic(err.Error())
	}
	DB = gstore.InitDB(dc.MaxIdleConn, dc.MaxOpenConn, dc.ConnMaxLifetime, dc.Dsn, glogs.Default(glogs.Sugar, logger.Info))

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

	// kafka消费端和生成端
	gkc, err := gconf.GetKafkaConf(kafkaKey)
	if err != nil {
		panic(err.Error())
	}
	var kc = &gkafka.KafkaConfig{
		Brokers:      gkc.Brokers,
		Topic:        gkc.Topic,
		Group:        gkc.Group,
		User:         gkc.User,
		Pwd:          gkc.Pwd,
		CertFile:     gkc.CertFile,
		KeyFile:      gkc.KeyFile,
		CaFile:       gkc.CaFile,
		KafkaVersion: gkc.KafkaVersion,
		VerifySsl:    gkc.VerifySsl,
	}
	Kafka = gkafka.InitKafka(kc)
}
