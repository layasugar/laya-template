//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya/store/db"
	"github.com/layasugar/laya/store/edb"
	"github.com/layasugar/laya/store/mdb"
	"github.com/layasugar/laya/store/rdb"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Orm orm
func Orm(ctx *laya.Context, dbName ...string) *gorm.DB {
	db := db.Wrap(ctx, dbName...)
	return db.WithContext(ctx)
}

// Rdb redis 连接
func Rdb(dbName ...string) *redis.Client {
	return rdb.GetClient(dbName...)
}

func Mdb(dbName ...string) *mongo.Client {
	return mdb.GetClient(dbName...)
}

func Edb(dbName ...string) *elastic.Client {
	return edb.GetClient(dbName...)
}
