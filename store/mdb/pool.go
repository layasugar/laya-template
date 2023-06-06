// mdbx 不考虑并发读写问题, 项目初始化就会初始化连接, 后面只会取
// map并发读没有问题

package mdb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	defaultMdbName = "default"
)

var p map[string]*mongo.Client

func getMdb(databaseName string) *mongo.Client {
	pool := getPool()
	return pool[databaseName]
}

func setMdb(databaseName string, db *mongo.Client) {
	pool := getPool()
	pool[databaseName] = db
}

func getPool() map[string]*mongo.Client {
	if p == nil {
		p = make(map[string]*mongo.Client)
	}
	return p
}
