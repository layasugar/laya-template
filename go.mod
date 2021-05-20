module github.com/layatips/laya-go

go 1.15

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis/v8 v8.8.0
	github.com/layatips/laya v0.0.8
	github.com/olivere/elastic/v6 v6.2.1
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.5.1
	gorm.io/gorm v1.21.7
)

//replace github.com/layatips/laya => ./../laya
