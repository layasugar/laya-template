package mdb

import (
	"context"
	"log"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var defaultPoolMaxOpen = uint64(runtime.NumCPU()*2 + 5) // 连接池最大连接数量4c*2+4只读副本+1主实例

const (
	defaultPoolMaxIdle     = 2                                // 连接池空闲连接数量
	defaultConnMaxIdleTime = time.Second * time.Duration(600) // 设置连接10分钟没有用到就断开连接(内存要求较高可降低该值)
)

// dbConfig Cluser Base Config
type dbConfig struct {
	name            string
	dsn             string
	poolMaxOpen     uint64
	poolMaxIdle     uint64
	connMaxIdleTime time.Duration
}

func InitConn(m []map[string]interface{}) {
	for _, item := range m {
		var dbc = dbConfig{}

		if name, ok := item["name"]; ok {
			if nameStr, okInterface := name.(string); okInterface {
				if nameStr == "" {
					dbc.name = defaultMdbName
				} else {
					dbc.name = nameStr
				}
			}
		} else {
			dbc.name = defaultMdbName
		}

		if dsn, ok := item["dsn"]; ok {
			if dsnStr, okInterface := dsn.(string); okInterface {
				dbc.dsn = dsnStr
			}
		}

		if poolMaxOpen, ok := item["max_open_conn"]; ok {
			if poolMaxOpenInt, okInterface := poolMaxOpen.(int64); okInterface {
				if poolMaxOpenInt == 0 {
					dbc.poolMaxOpen = defaultPoolMaxOpen
				} else {
					dbc.poolMaxOpen = uint64(poolMaxOpenInt)
				}
			}
		}

		if poolMaxIdle, ok := item["max_idle_conn"]; ok {
			if poolMaxIdleInt, okInterface := poolMaxIdle.(int64); okInterface {
				if poolMaxIdleInt == 0 {
					dbc.poolMaxIdle = defaultPoolMaxIdle
				} else {
					dbc.poolMaxIdle = uint64(poolMaxIdleInt)
				}
			}
		}

		if connMaxIdleTime, ok := item["max_idle_time"]; ok {
			if connMaxIdleTimeInt, okInterface := connMaxIdleTime.(int64); okInterface {
				if connMaxIdleTimeInt == 0 {
					dbc.connMaxIdleTime = defaultConnMaxIdleTime
				} else {
					dbc.connMaxIdleTime = time.Second * time.Duration(connMaxIdleTimeInt)
				}
			}
		}

		setMdb(dbc.name, dbc.Open())
	}
}

// Open provides method to opening a database link with DBConfig struct
func (c *dbConfig) Open() *mongo.Client {
	var t = NewTracer()
	cmdMonitor := &event.CommandMonitor{
		Started:   t.HandleStartedEvent,
		Succeeded: t.HandleSucceededEvent,
		Failed:    t.HandleFailedEvent,
	}

	opts := &options.ClientOptions{}
	opts.ApplyURI(c.dsn).
		SetMaxConnecting(c.poolMaxOpen).
		SetMinPoolSize(c.poolMaxIdle).
		SetMaxConnIdleTime(c.connMaxIdleTime).
		SetMonitor(cmdMonitor)
	con, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	err = con.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	log.Printf("[app.mdb] mongo success, name: %s", c.name)
	return con
}

func GetClient(name ...string) *mongo.Client {
	if len(name) > 0 {
		return getMdb(name[0])
	} else {
		return getMdb(defaultMdbName)
	}
}
