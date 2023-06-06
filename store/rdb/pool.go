// rdbx 不考虑并发读写问题, 项目初始化就会初始化连接, 后面只会取
// map并发读没有问题

package rdb

import (
	"github.com/go-redis/redis/v8"
)

const (
	defaultRdbName = "default"
)

var p map[string]*redis.Client

func getRdb(name string) *redis.Client {
	pool := getPool()
	return pool[name]
}

func setRdb(databaseName string, db *redis.Client) {
	pool := getPool()
	pool[databaseName] = db
}

func getPool() map[string]*redis.Client {
	if p == nil {
		p = make(map[string]*redis.Client)
	}
	return p
}
