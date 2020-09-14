package ship

import (
	"encoding/json"
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
	"io/ioutil"
	"time"
)

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

func Init() {
	InitEnv()
	InitMysql()
	InitRedis()
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

	// get i18n config
	err = json.Unmarshal(config.Get(ENV, "i18n").Bytes(), &I18nConf)
	if err != nil {
		panic(err)
	}

	// get mysql config
	err = json.Unmarshal(config.Get(ENV, "database").Bytes(), &MysqlConf)
	if err != nil {
		panic(err)
	}

	// get cache config
	err = json.Unmarshal(config.Get(ENV, "cache").Bytes(), &RedisConf)
	if err != nil {
		panic(err)
	}

	// get delayerServer config
	DelayServer = config.Get(ENV, "delayServer").String("http://127.0.0.1:9278")
}

// 初始化redis
func InitRedis() {
	if RedisConf.Open {
		Redis = redis.NewClient(&redis.Options{
			Addr:        RedisConf.Addr,                                     // Redis地址
			Password:    RedisConf.Pwd,                                      // Redis账号
			DB:          RedisConf.DB,                                       // Redis库
			PoolSize:    RedisConf.PoolSize,                                 // Redis连接池大小
			MaxRetries:  RedisConf.MaxRetries,                               // 最大重试次数
			IdleTimeout: time.Second * time.Duration(RedisConf.IdleTimeout), // 空闲链接超时时间
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
}

// 初始化mysql
func InitMysql() {
	if MysqlConf.Open {
		var err error
		DB, err = gorm.Open("mysql", MysqlConf.Dsn)
		if err != nil {
			panic(err)
		}
		DB.DB().SetMaxIdleConns(MysqlConf.MaxIdleConn)
		DB.DB().SetMaxOpenConns(MysqlConf.MaxOpenConn)
		DB.DB().SetConnMaxLifetime(time.Hour * time.Duration(MysqlConf.ConnMaxLifetime))
	}
}

func InitLang() {
	if I18nConf.Open {
		I18nBundle = i18n.NewBundle(language.English)
		I18nBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		err := LoadAllFile("./lang/")
		if err != nil {
			panic(err)
		}
	}
}

func LoadAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			_ = LoadAllFile(pathname + fi.Name() + "\\")
		} else {
			_, err := I18nBundle.LoadMessageFile(pathname + fi.Name())
			if err != nil {
				return err
			}
		}
	}
	return err
}
