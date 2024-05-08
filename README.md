# laya-template

- http框架模板, grpc框架模板, task服务模板
- 快速搭建中小型应用服务, restfulApi, grpc服务，web服务

> Please note that this repository is out-of-the-box template

## 使用之前

- 这不是一个框架，是个人为简单编码整理的项目（服务）目录结构和代码结构
- 适合中小型项目和应用程序
- 不支持swagger，接口文档需要编码人员精心编写，代码保持干净整洁
- 配置文件不支持热重载，更新配置文件需要重启服务，或者自己实现
- 简单，快速，高效

## 约定

- func返回单独结构体时, 返回该数据的指针
- core.Context需要全局传递(ctx里面内置了记录日志与链路追踪)
- models/page 业务逻辑
- models/data 实现数据查询组装, 查询在此处完成, 尽量不要使用join(减轻数据库压力), 数据取出后, 可在该层完成组装
- models/dao 基本的请求层, 模型放置层
- 非必要无须grpc连接池

## 使用
```shell
git clone git@github.com:layasugar/laya-template.git

cd laya-template && go mod tidy

go run .
```

## 感谢以下开源仓库
- [gin](https://github.com/gin-gonic/gin)
- [gorm文档](https://gorm.io/zh_CN/docs/index.html)
- [go-redis](https://redis.uptrace.dev/)
- [mongo文档](https://www.mongodb.com/docs/drivers/go/current/usage-examples/)
- [es文档](https://olivere.github.io/elastic/)

### License

`laya-template` is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.

### 🔑 JetBrains OS licenses

`laya-template` had been being developed with `GoLand` IDE under the **free JetBrains Open Source license(s)** granted by JetBrains s.r.o., hence I would like to express my thanks here.

<a href="https://www.jetbrains.com/?from=gnet" target="_blank"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" width="250" align="middle"/></a>