//数据库连接基础文件，根据自己需要定制

package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/laya/gstore/dbx"
	"gorm.io/gorm"
)

// DB is sql *db
var DB *gorm.DB

// Rdb is redis *client
var Rdb *redis.Client

func Init() {
	//// mysql
	//DB = gstore.InitDB(gconf.V.GetString("mysql.dsn"), gstore.LevelInfo)
	//
	//// redis
	//rdbCfg := redis.Options{
	//	Addr:     gconf.V.GetString("redis.addr"),
	//	DB:       gconf.V.GetInt("redis.db"),
	//	Password: gconf.V.GetString("redis.pwd"),
	//}
	//Rdb = gstore.InitRdb(rdbCfg)
}

// Orm orm
func Orm(ctx context.Context, dbName ...string) *gorm.DB {
	db := dbx.Wrap(ctx, dbName...)
	return db.WithContext(ctx)
}
