package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layatips/laya/gconf"
	"github.com/layatips/laya/gstore"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// DB is sql *db
var DB *gorm.DB
var dbKey = "mysql"

// Rdb is redis *client
var Rdb *redis.Client
var rdbKey = "redis"

// Mdb is mongo client
var Mdb *mongo.Client
var mdbKey = "mongo"

func InitDao() {
	dbConfig, err := gconf.GetDBConf(dbKey)
	if err != nil {
		panic("err db config")
	}
	DB = gstore.InitDB(dbConfig)

	rdbConfig, err := gconf.GetRdbConf(rdbKey)
	if err != nil {
		panic("err rdb config")
	}
	Rdb = gstore.InitRdb(rdbConfig)

	mdbConfig, err := gconf.GetMdbConf(mdbKey)
	Mdb = gstore.InitMdb(mdbConfig)
}
