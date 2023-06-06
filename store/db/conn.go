package db

import (
	"log"
	"runtime"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var defaultPoolMaxOpen = runtime.NumCPU()*2 + 5 // 连接池最大连接数量4c*2+4只读副本+1主实例

const (
	defaultPoolMaxIdle     = 2                                 // 连接池空闲连接数量
	defaultConnMaxLifeTime = time.Second * time.Duration(7200) // MySQL默认长连接时间为8个小时,可根据高并发业务持续时间合理设置该值
	defaultConnMaxIdleTime = time.Second * time.Duration(60)   // 设置连接10分钟没有用到就断开连接(内存要求较高可降低该值)
	levelInfo              = "info"
	levelWarn              = "warn"
	levelError             = "error"
)

type dbPoolCfg struct {
	maxIdleConn int64 // 空闲连接数
	maxOpenConn int64 // 最大连接数
	maxLifeTime int64 // 连接可重用的最大时间
	maxIdleTime int64 // 在关闭连接之前, 连接可能处于空闲状态的最大时间
}

type dbConfig struct {
	name     string
	dsn      string
	logLevel string
	poolCfg  *dbPoolCfg
	gormCfg  *gorm.Config
}

// initDB init db
func initDB(cfg dbConfig) *gorm.DB {
	var err error

	var level logger.LogLevel
	switch cfg.logLevel {
	case levelInfo:
		level = logger.Info
	case levelWarn:
		level = logger.Warn
	case levelError:
		level = logger.Error
	default:
		level = logger.Info
	}

	cfg.gormCfg = &gorm.Config{
		Logger: Default(level),
	}

	Db, err := gorm.Open(mysql.Open(cfg.dsn), cfg.gormCfg)
	if err != nil {
		log.Printf("[app.db] mysql open fail, err:%s", err)
		panic(err)
	}

	cfg.setDefaultPoolConfig(Db)

	registerCallbacks(Db)

	err = DbSurvive(Db)
	if err != nil {
		log.Printf("[app.db] mysql survive fail, err:%s", err)
		panic(err)
	}

	log.Printf("[app.db] mysql success, name: %s", cfg.name)
	return Db
}

// DbSurvive mysql survive
func DbSurvive(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (c *dbConfig) setDefaultPoolConfig(db *gorm.DB) {
	d, err := db.DB()
	if err != nil {
		log.Printf("[app.dbx] mysql db fail, err: %s", err.Error())
		panic(err)
	}
	var cfg = c.poolCfg
	if cfg == nil {
		d.SetMaxOpenConns(defaultPoolMaxOpen)
		d.SetMaxIdleConns(defaultPoolMaxIdle)
		d.SetConnMaxLifetime(defaultConnMaxLifeTime)
		d.SetConnMaxIdleTime(defaultConnMaxIdleTime)
		return
	}

	if cfg.maxOpenConn == 0 {
		d.SetMaxOpenConns(defaultPoolMaxOpen)
	} else {
		d.SetMaxOpenConns(int(cfg.maxOpenConn))
	}

	if cfg.maxIdleConn == 0 {
		d.SetMaxIdleConns(defaultPoolMaxIdle)
	} else {
		d.SetMaxIdleConns(int(cfg.maxIdleConn))
	}

	if cfg.maxLifeTime == 0 {
		d.SetConnMaxLifetime(defaultConnMaxLifeTime)
	} else {
		d.SetConnMaxLifetime(time.Second * time.Duration(cfg.maxLifeTime))
	}

	if cfg.maxIdleTime == 0 {
		d.SetConnMaxIdleTime(defaultConnMaxIdleTime)
	} else {
		d.SetConnMaxIdleTime(time.Second * time.Duration(cfg.maxIdleTime))
	}
}

func InitConn(m []map[string]interface{}) {
	for _, item := range m {
		var dbc = dbConfig{}
		var poolCfg dbPoolCfg

		if name, ok := item["name"]; ok {
			if nameStr, okInterface := name.(string); okInterface {
				if nameStr == "" {
					dbc.name = defaultDbName
				} else {
					dbc.name = nameStr
				}
			}
		} else {
			dbc.name = defaultDbName
		}

		if dsn, ok := item["dsn"]; ok {
			if dsnStr, okInterface := dsn.(string); okInterface {
				dbc.dsn = dsnStr
			}
		}

		if level, ok := item["level"]; ok {
			if levelStr, okInterface := level.(string); okInterface {
				dbc.logLevel = levelStr
			}
		}

		if maxIdleConn, ok := item["max_idle_conn"]; ok {
			if maxIdleConnInt, okInterface := maxIdleConn.(int64); okInterface {
				poolCfg.maxIdleConn = maxIdleConnInt
			}
		}

		if maxOpenConn, ok := item["max_open_conn"]; ok {
			if maxOpenConnInt, okInterface := maxOpenConn.(int64); okInterface {
				poolCfg.maxOpenConn = maxOpenConnInt
			}
		}

		if maxLifeTime, ok := item["max_life_time"]; ok {
			if maxLifeTimeInt, okInterface := maxLifeTime.(int64); okInterface {
				poolCfg.maxLifeTime = maxLifeTimeInt
			}
		}

		if maxIdleTime, ok := item["max_idle_time"]; ok {
			if maxIdleTimeInt, okInterface := maxIdleTime.(int64); okInterface {
				poolCfg.maxIdleTime = maxIdleTimeInt
			}
		}

		if dbc.name == "" || dbc.dsn == "" {
			continue
		}
		dbc.poolCfg = &poolCfg

		setGormDB(dbc.name, initDB(dbc))
	}
}
