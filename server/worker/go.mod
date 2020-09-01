module laya-go/server/worker

go 1.13

require (
	laya-go/ship v1.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.1.0
	github.com/robfig/cron/v3 v3.0.0
)

replace laya-go/ship => ./../../ship
