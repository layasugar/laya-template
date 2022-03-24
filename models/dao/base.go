//数据库连接基础文件，根据自己需要定制

package dao

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/laya/store/dbx"
	"github.com/layasugar/laya/store/edbx"
	"github.com/layasugar/laya/store/mdbx"
	"github.com/layasugar/laya/store/rdbx"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Orm orm
func Orm(ctx context.Context, dbName ...string) *gorm.DB {
	db := dbx.Wrap(ctx, dbName...)
	return db.WithContext(ctx)
}

// Rdb redis 连接
func Rdb(dbName ...string) *redis.Client {
	return rdbx.GetClient(dbName...)
}

func Mdb(dbName ...string) *mongo.Client {
	return mdbx.GetClient(dbName...)
}

func Edb(dbName ...string) *elasticsearch.Client {
	return edbx.GetClient(dbName...)
}
