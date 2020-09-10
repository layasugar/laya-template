package ship

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)
// environment
var ENV string

// config value
var MysqlDsn string
var RedisAddr string
var RedisPwd string
var DefaultLang string
var DelayServer string

// 定义redis链接池,mysql连接池,语言包bundle
var Redis *redis.Client
var DB *gorm.DB
var I18nBundle *i18n.Bundle
