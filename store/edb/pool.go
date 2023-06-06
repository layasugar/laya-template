// edbx 不考虑并发读写问题, 项目初始化就会初始化连接, 后面只会取
// map并发读没有问题

package edb

import (
	"github.com/olivere/elastic/v7"
)

const (
	defaultEdbName = "default"
)

var p map[string]*elastic.Client

func getEdb(databaseName string) *elastic.Client {
	pool := getPool()
	return pool[databaseName]
}

func setEdb(databaseName string, db *elastic.Client) {
	pool := getPool()
	pool[databaseName] = db
}

func getPool() map[string]*elastic.Client {
	if p == nil {
		p = make(map[string]*elastic.Client)
	}
	return p
}
