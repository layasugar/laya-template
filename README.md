# laya-go
gin+gorm+go-redis为基本骨架打造的框架，方便开发，开箱即用，micro和etcd可选用，针对大项目微服务架构

## i18n
[详细介绍](https://github.com/LaYa-op/laya-go/wiki/i18n)

- 简单使用
```
c.Set("$.TokenErr.code", r.TokenErr)
会用"."隔开，取出倒数第二段字符TokenErr作为ID，去翻译文件取出翻译
```

## 配置文件的使用


## 接口签名

## 延时队列

## 网关

## 注册中心

## 中间键

## 文件上传和处理

## redis锁

## 运行配置
```
--registry=etcdv3
```

## 网关启动
```
.\micro.exe --registry=etcdv3 api --handler=proxy --namespace=tb.server --address=:10081
```