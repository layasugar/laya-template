module github.com/LaYa-op/laya-go

go 1.15

require (
	github.com/LaYa-op/laya v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/nicksnyder/go-i18n/v2 v2.1.1
	golang.org/x/text v0.3.4
)

replace (
	github.com/LaYa-op/laya => ./../laya
)
