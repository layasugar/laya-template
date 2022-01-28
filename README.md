# laya-template

后端服务模板 使用 [laya](https://github.com/layasugar/laya) 搭建 旨在快速搭建中小型应用服务, restfulApi, rpc服务 ==

> Please note that this repository is out-of-the-box template

## 约定

- func返回单独结构体时, 返回该数据得指针
- laya.WebContext与laya.PbRpcContext需要全局传递
- models/page 实现业务
- models/data 实现数据查询组装, 查询在此处完成, 不要使用join, 数据取出后, 可在该层完成组装
- models/dao 基本的请求层, 模型放置层

## 快速使用

### web, grpc应用

- 参考mian.go

## 推荐工具

### 数据库直接生成gorm的struct

- [github链接](https://github.com/Shelnutt2/db2struct)
- ```db2struct --host localhost -d database --package db -p 123456 --user root --guregu --gorm -t tableName --struct structName```

### hey 压测工具

- [github链接](https://github.com/rakyll/hey)
- ```hey -n 100 -c 1000```

### gorm外封装一层, 处理日志, 自动将gorm日志定向到logger