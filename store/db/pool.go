// dbx 不考虑并发读写问题, 项目初始化就会初始化连接, 后面只会取
// map并发读没有问题

package db

import "gorm.io/gorm"

const (
	defaultDbName = "default"
)

var dbPool map[string]*gorm.DB

func getGormDB(databaseName string) *gorm.DB {
	pool := getPool()
	return pool[databaseName]
}

func setGormDB(databaseName string, db *gorm.DB) {
	pool := getPool()
	pool[databaseName] = db
}

func getPool() map[string]*gorm.DB {
	if dbPool == nil {
		dbPool = make(map[string]*gorm.DB)
	}
	return dbPool
}
