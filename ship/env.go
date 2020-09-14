package ship

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// environment
var ENV string

var DelayServer string

var RedisConf struct {
	Open        bool   `json:"open"`
	DB          int    `json:"db"`
	Addr        string `json:"addr"`
	Pwd         string `json:"pwd"`
	PoolSize    int    `json:"poolSize"`
	MaxRetries  int    `json:"maxRetries"`
	IdleTimeout int    `json:"idleTimeout"`
}

var MysqlConf struct {
	Open            bool   `json:"open"`
	Dsn             string `json:"dsn"`
	MaxIdleConn     int    `json:"maxIdleConn"`
	MaxOpenConn     int    `json:"maxOpenConn"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
}

var I18nConf struct {
	Open        bool   `json:"open"`
	DefaultLang string `json:"defaultLang"`
}

// 定义redis链接池,mysql连接池,语言包bundle
var Redis *redis.Client
var DB *gorm.DB
var I18nBundle *i18n.Bundle
