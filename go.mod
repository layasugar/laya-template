module github.com/layatips/laya-go

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.4.0
	github.com/layatips/laya v0.0.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	gorm.io/gorm v1.20.7
)

replace github.com/layatips/laya => ./../laya
