module laya-go/server/hall

go 1.13

require (
	//laya-go/server/common/model v1.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.1.0
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.2
	laya-go/common v1.0.0
	laya-go/ship v1.0.0
)

replace laya-go/ship => ./../../ship

replace laya-go/common => ./../../common
