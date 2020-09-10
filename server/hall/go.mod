module laya-go/server/hall

go 1.13

require (
	github.com/clevergo/captchas v0.3.2
	github.com/clevergo/captchas/drivers v0.3.2
	github.com/clevergo/captchas/stores/redisstore v0.1.2
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v7 v7.3.0
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.2-0.20200203150404-894195727d9c
	github.com/micro/go-micro/v2 v2.1.0
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.2
	laya-go/ship v1.0.0
)

replace laya-go/ship => ./../../ship
