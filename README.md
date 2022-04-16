# laya-template

- http框架模板, grpc框架模板, 默认服务模板
- 使用 [laya](https://github.com/layasugar/laya) 搭建 旨在快速搭建中小型应用服务, restfulApi, rpc服务 ==

> Please note that this repository is out-of-the-box template

## 约定

- func返回单独结构体时, 返回该数据得指针
- laya.WebContext与laya.GrpcContext需要全局传递(ctx里面内置了记录日志与链路追踪)
- models/page 业务逻辑
- models/data 实现数据查询组装, 查询在此处完成, 尽量不要使用join(减轻数据库压力), 数据取出后, 可在该层完成组装
- models/dao 基本的请求层, 模型放置层

## 安装模板, 愉快编码
#### 安装
```shell
go install github.com/layasugar/laya/laya@latest
```
#### 初始化模板
```shell
laya template init -name=laya-template

laya template init-http -name=laya-template

laya template init-grpc -name=laya-template
```
#### 运行
```shell
cd xxx;
go get -u github.com/layasugar/laya
go mod tidy
go run .
```

## laya-template 体验
- 该模板本身就是一个demo, 覆盖laya所有功能的测试
- [体验操作文档](https://github.com/layasugar/laya-template/blob/master/experience.md)
- [gorm文档](https://gorm.io/zh_CN/docs/index.html)
- [go-redis](https://redis.uptrace.dev/)
- [mongo文档](https://www.mongodb.com/docs/drivers/go/current/usage-examples/)
- [es文档](https://olivere.github.io/elastic/)