//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/laya/gcache"
	"github.com/layasugar/laya/gconf"
	"github.com/layasugar/laya/gkafka"
	"github.com/layasugar/laya/glogs"
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

	// es客户端
	esc, err := gconf.GetEsConf(esKey)
	if err != nil {
		panic(err.Error())
	}
	Es, err = elastic.NewClient(
		elastic.SetURL(esc.Addr...),
		elastic.SetBasicAuth(esc.User, esc.Pwd),
	)
	if err != nil {
		panic(err.Error())
	}
}
