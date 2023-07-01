//数据库连接基础文件，根据自己需要定制

package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/layasugar/laya-template/pkg/core"
	"github.com/layasugar/laya-template/store/db"
	"github.com/layasugar/laya-template/store/rdb"
	"gorm.io/gorm"
)

// Orm mysql 连接池
func Orm(ctx *core.Context, dbName ...string) *gorm.DB {
	conn := db.Wrap(ctx, dbName...)
	return conn.WithContext(ctx)
}

// Rdb redis 连接池
func Rdb(dbName ...string) *redis.Client {
	return rdb.GetClient(dbName...)
}
