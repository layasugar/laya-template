module github.com/layasugar/laya-go

go 1.15

require (
	github.com/gin-gonic/gin v1.7.3
	github.com/go-redis/redis/v8 v8.11.2
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/layasugar/laya v0.1.0
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/olivere/elastic/v7 v7.0.27
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.7.1
	gorm.io/gorm v1.21.12
)

replace github.com/layasugar/laya => ./../laya
