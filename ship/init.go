package ship

import (
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"time"
)

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

func Before() {
	app := cmd.App()
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:  "env",
		Usage: "environment to setting",
	})

	before := app.Before
	app.Before = func(ctx *cli.Context) error {
		if path := ctx.String("env"); len(path) > 0 {
			// got config
			// do stuff
			ENV = path
		} else {
			ENV = DefaultEnv
		}
		return before(ctx)
	}
}

func Init(rc RedisConf, mc MysqlConf) {
	InitEnv()
	if rc.Open {
		InitRedis(RedisAddr, RedisPwd, rc)
	}

	if mc.Open {
		InitMysql(MysqlDsn, mc)
	}
	InitLang()
}

func InitEnv() {
	_, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	err = config.LoadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	MysqlDsn = config.Get(ENV, "database", "dsn").String("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	RedisAddr = config.Get(ENV, "cache", "addr").String("127.0.0.1:6379")
	RedisPwd = config.Get(ENV, "cache", "password").String("123456")
	DefaultLang = config.Get(ENV, "defaultLang").String("zh")
	DelayServer = config.Get(ENV, "delayServer").String("http://127.0.0.1:9278")
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

func InitLang() {
	I18nBundle = i18n.NewBundle(language.English)
	I18nBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := I18nBundle.LoadMessageFile("./lang/zh.toml")
	if err != nil {
		panic(err)
	}
	_, err = I18nBundle.LoadMessageFile("./lang/en.toml")
	if err != nil {
		panic(err)
	}
}
