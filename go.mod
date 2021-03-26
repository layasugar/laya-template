module github.com/layatips/laya-go

go 1.15

require (
	github.com/layatips/laya v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.4.0
	gorm.io/gorm v1.20.7
)

replace github.com/layatips/laya => ./../laya
