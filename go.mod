module github.com/layatips/laya-go

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.4.0
	github.com/layatips/laya v0.0.6
	github.com/spf13/viper v1.7.1 // indirect
	go.mongodb.org/mongo-driver v1.4.3
	gorm.io/gorm v1.20.7
)

replace github.com/layatips/laya => ./../laya
