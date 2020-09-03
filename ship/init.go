package ship

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"time"
)

// 定义redis链接池,mysql连接池
var Redis *redis.Client
var DB *gorm.DB

type RedisConf struct {
	Open        bool
	DB          int
	PoolSize    int
	MaxRetries  int
	IdleTimeout time.Duration
}

type MysqlConf struct {
	Open            bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

func Init(rc RedisConf, mc MysqlConf) {
	_, _ = config.NewConfig()
	err := config.LoadFile("config.yaml")
	if err != nil {
		log.Info(err)
	}
	mysqlDsn := config.Get("database", "dsn").String("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	redisAddr := config.Get("cache", "addr").String("127.0.0.1:6379")
	redisPwd := config.Get("cache", "password").String("123456")

	if rc.Open {
		InitRedis(redisAddr, redisPwd, rc)
	}

	if mc.Open {
		InitMysql(mysqlDsn, mc)
	}
}

// 初始化redis
func InitRedis(addr string, pwd string, rc RedisConf) {
	Redis = redis.NewClient(&redis.Options{
		Addr:        addr,           // Redis地址
		Password:    pwd,            // Redis账号
		DB:          rc.DB,          // Redis库
		PoolSize:    rc.PoolSize,    // Redis连接池大小
		MaxRetries:  rc.MaxRetries,  // 最大重试次数
		IdleTimeout: rc.IdleTimeout, // 空闲链接超时时间
	})
	pong, err := Redis.Ping().Result()
	if err == redis.Nil {
		log.Info("Nil reply returned by Redis when key does not exist.")
	} else if err != nil {
		panic(err)
	} else {
		log.Info(pong)
	}
}

// 初始化mysql
func InitMysql(dsn string, mc MysqlConf) {
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(mc.MaxIdleConn)
	DB.DB().SetMaxOpenConns(mc.MaxOpenConn)
	DB.DB().SetConnMaxLifetime(mc.ConnMaxLifetime)
}
