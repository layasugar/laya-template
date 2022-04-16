## 体验

#### clone代码

```
git clone git@github.com:layasugar/laya-template.git
```

<hr>

#### 存储配置(mysql,redis,mongo,es), 参照conf/app.toml

- mysql 的表结构在script/test.sql, 创建laya_template数据库, 导入sql文件
- 修改redis 连接配置
- 修改mongo 连接配置
- 修改es 连接配置

##### 提供docker安装mysql, redis, mongo, es, zipkin, jaeger
```
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql --default-authentication-plugin=mysql_native_password

docker run -d -p 6379:6379 --name redis redis --requirepass "123456"

docker run -d --name mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 27017:27017 mongo

docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.17.1

docker run -d -p 9411:9411 --name zipkin openzipkin/zipkin
```

<hr>

#### 测试实例one

```
- 切换main函数启动一个http server
- POST JSON
{
    "kind": 1
}
- http://127.0.0.1:80/full-test
```

##### kind = 1 空转测试并发能力

- [x] 接口请求日志打印
- [x] 接口处理链路上报zipkin或者jaeger
- [x] pprof持续观测结果
- [x] 接口日志开关
- [x] 接口链路开关

##### kind = 2 测试mysql

- [x] mysql的curd的正确性
- [x] mysql多连接的正确性
- [x] mysql执行sql日志打印
- [x] mysql链路上报zipkin或者jaeger
- [x] pprof持续观测结果
- [x] mysql日志开关
- [x] mysql链路开关

##### kind = 3 测试redis

- [x] redis的crd正确性
- [x] redis多连接池的正确性
- [x] redis链路上报zipkin或者jaeger
- [x] redis链路上报开关

##### kind = 4 测试mongo .Database("test").Collection("user")

- [x] mongo的curd正确性
- [x] mongo多连接池的正确性
- [x] mongo链路上报zipkin或者jaeger
- [x] mongo链路上报开关

##### kind = 5 测试es

- [x] es的curd正确性
- [x] es多连接池的正确性
- [x] es链路上报zipkin或者jaeger
- [x] es链路上报开关

##### kind = 6 测试http_to_http

- [x] 服务日志打印
- [x] 服务处理链路上报zipkin或者jaeger

##### kind = 7 测试http_to_grpc

- [x] 服务日志打印
- [x] 服务处理链路上报zipkin或者jaeger

<hr>

#### 测试实例two

```
- 切换main函数启动一个grpc server
- GRPC DATA
{
    "kind": 1
}
- grpc://127.0.0.1:80/SayHello
- grpc://127.0.0.1:80/GrpcTraceTest
```

- [x] kind = 1 空转grpc
- [x] kind = 2 测试grpc_to_http
- [x] kind = 2 测试grpc_to_grpc

<hr>

#### 测试实例three

```
- 切换main函数启动一个app server
NODATA
```

- [x] kind = 1 测试app_to_http
- [x] kind = 2 测试app_to_grpc

<hr>