module laya-go/server/hall

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.1.0
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.2
	laya-go/ship v1.0.0
)

replace laya-go/ship => ./../../ship