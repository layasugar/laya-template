//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layatips/laya/gcache"
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

// Mem
var Mem *gcache.Cache

func InitDao() {
	//mysql
	dbConfig, err := gconf.GetDBConf(dbKey)
	if err != nil {
		panic(err.Error())
	}
	DB = gstore.InitDB(dbConfig)

	//redis
	rdbConfig, err := gconf.GetRdbConf(rdbKey)
	if err != nil {
		panic(err.Error())
	}
	Rdb = gstore.InitRdb(rdbConfig)

	//mongodb
	mdbConfig, err := gconf.GetMdbConf(mdbKey)
	if err != nil {
		panic(err.Error())
	}
	Mdb = gstore.InitMdb(mdbConfig)

	//mem
	Mem = gstore.InitMemory()
}
