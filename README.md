# laya-go
gin+gorm+go-redis为基本骨架打造的框架，方便开发，开箱即用

i18n
------------
[i18n](https://github.com/LaYa-op/laya-go/wiki/i18n)



接口签名

### 延时队列

### 网关

### 注册中心

### 中间键

### 文件上传和处理

### redis锁

###运行配置：
```
--registry=etcdv3
```

###网关启动：
```
.\micro.exe --registry=etcdv3 api --handler=proxy --namespace=tb.server --address=:10081
```