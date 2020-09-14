# laya-go
gin+gorm+go-redis为基本骨架打造的框架，方便开发，开箱即用，micro和etcd可选用，针对大项目微服务架构

## i18n
[详细介绍](https://github.com/LaYa-op/laya-go/wiki/i18n)

- [go-i18n](https://github.com/nicksnyder/go-i18n)
- 简单使用
```
c.Set("$.TokenErr.code", r.TokenErr)
会用"."隔开，取出倒数第二段字符TokenErr作为ID翻译
```

## 配置文件的使用
- 运行服务时，请添加启动参数--env=dev,表示使用dev配置，具体配置请参照config.yaml

## 接口签名
- 接口签名是按照gin标准中间件编写的，你可以在采用一下方式使用
```
r.Use(middleware.Sign())
```
- 或者在路由分组中使用

## 延时队列
- [delay-queue](https://github.com/ouqiang/delay-queue)
- ship/utils/job.go下提供4个方法JobPush和JobPop和JobFinish和JobRemove
- 简单使用
```
ship.JobPush
ship.JobPop
ship.JobFinish
ship.JobRemove
```

## 网关
- [micro](https://github.com/micro/micro)
- 简单启动
```
micro api --registry=etcdv3  --handler=http --namespace=laya-go.server --address=:10081
```

## 注册中心
- [etcd](https://github.com/etcd-io/etcd)
- 单机启动
```
etcd
```
- 集群启动
```
etcd --
```

## 中间键
- [gin-middleware](https://github.com/gin-gonic/gin)

## 文件上传和处理

## redis锁
- ship/utils/redis_lock.go下提供2个方法GetLock和ReleaseLock