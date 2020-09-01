package config

import "time"

// 系统配置
const (
    OpenRedis = true // 开启redis
    OpenMysql = true // 开启mysql

    // redis配置
    RedisDB          = 1                // Redis库
    RedisPoolSize    = 5                // Redis连接池大小
    RedisMaxRetries  = 3                // 最大重试次数
    RedisIdleTimeout = 10 * time.Second // 空闲链接超时时间

    // mysql配置
    //MaxIdleConn     = 5  // mysql最大空闲链接数
    //MaxOpenConn     = 50 // mysql最大连接数
    //ConnMaxLifetime = time.Second * 10
    MaxIdleConn     = 10  // mysql最大空闲链接数
    MaxOpenConn     = 100 // mysql最大连接数
    ConnMaxLifetime = time.Hour * 6
)
