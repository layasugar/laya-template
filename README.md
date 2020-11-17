# laya-go
gin+gorm+go-redis为基本骨架打造的框架，方便开发，开箱即用，micro和etcd可选用，针对中型项目微服务架构

[目录结构介绍](https://github.com/LaYa-op/laya-go/wiki/%E7%9B%AE%E5%BD%95%E7%BB%93%E6%9E%84%E4%BB%8B%E7%BB%8D)
[单个服务启动](https://github.com/LaYa-op/laya-go/wiki/%E5%8D%95%E4%B8%AA%E6%9C%8D%E5%8A%A1%E5%90%AF%E5%8A%A8)
[微服务方式启动](https://github.com/LaYa-op/laya-go/wiki/%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%96%B9%E5%BC%8F%E5%90%AF%E5%8A%A8)

## i18n
[i18n使用介绍](https://github.com/LaYa-op/laya-go/wiki/i18n)

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
