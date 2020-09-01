package main

import (
    "laya-go/ship"
    "laya-go/server/worker/config"
)

// redis配置
var RC = ship.RedisConf{
    Open:        config.OpenRedis,
    DB:          config.RedisDB,
    PoolSize:    config.RedisPoolSize,
    MaxRetries:  config.RedisMaxRetries,
    IdleTimeout: config.RedisIdleTimeout,
}

// mysql配置
var MC = ship.MysqlConf{
    Open:            config.OpenMysql,
    MaxIdleConn:     config.MaxIdleConn,
    MaxOpenConn:     config.MaxOpenConn,
    ConnMaxLifetime: config.ConnMaxLifetime,
}

func Init() {
    ship.Init(RC, MC)
}
