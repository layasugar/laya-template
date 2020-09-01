module laya-go/base

go 1.14

require (
	laya-go/ship v1.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v7 v7.3.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.1.0
	github.com/oschwald/geoip2-golang v1.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/thinkeridea/go-extend v1.1.1
)

replace laya-go/ship => ./../ship