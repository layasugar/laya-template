package dao

import (
	"github.com/layatips/laya/gstore"
)

// db is sql *db
var DB = gstore.DB

// rdb is redis *db
var Rdb = gstore.Rdb

// mdb is mongodb *mongo.DB
var Mdb = gstore.Mdb
