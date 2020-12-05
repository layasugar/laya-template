module github.com/LaYa-op/laya-go

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/LaYa-op/laya v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.9.1
	gorm.io/gorm v1.20.7
)

replace github.com/LaYa-op/laya => ./../laya
