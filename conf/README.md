### 配置文件说明

#### 项目app基础配置

```
  "app": {
    "app_name": "laya-go",
    "http_listen": "0.0.0.0:10080",
    "run_mode": "debug",
    "app_url": "https://github.com/layasugar/laya-go",
    "gin_log": "/home/logs/app/laya-go/gin_http.log",
    "params_log": true
  }
## app_name：项目名称
## http_listen：监听端口
## run_mode: debug or release
## app_url: 当前项目请求url
## gin_log: gin的http请求日志输出到文件
## params_log: 是否打印入参和出参
```

##

### log

#### https://github.com/uber-go/zap

```
  "log": {
    "path": "/home/logs/app/payment/app.log",
    "max_size": 32,
    "max_age": 90,
    "max_backups": 300
  }
## path：日志路径
## max_size：单个日志大小xxMB
## max_age：日志保留最大天数xx天
## max_backups：保留旧文件的最大个数xx个
```

##

### cache

#### https://github.com/patrickmn/go-cache

```
  "mem": {
    "default_exp": 300,
    "cleanup": 600
  }
## default_exp：秒-默认过期时间
## cleanup：秒-每多少清除一次过期物品
```

##

### mysql

#### http://gorm.io/zh_CN/docs/index.html

```
  "mysql": {
    "dsn": "root:123456@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local",
    "maxIdleConn": 10,
    "maxOpenConn": 100,
    "connMaxLifetime": 6
  }
## dsn：db DSN
## maxIdleConn：db连接池最大空闲链接数
## maxOpenConn：db连接池最大连接数
## connMaxLifetime：连接池超时时间(单位: time.Hour)
```

##

### redis

#### https://github.com/go-redis/redis

```
  "redis": {
    "db": 0,
    "addr": "redis:6379",
    "pwd": "",
    "poolSize": 5,
    "maxRetries": 3,
    "idleTimeout": 1000
  }
## db：db
## addr：链接DSN
## pwd：链接密码
## poolSize：链接池大小
## maxRetries：链接最大重试次数
## idleTimeout：空闲链接超时时间(单位：time.Second秒)
```

##       

### mongo

#### https://github.com/mongodb/mongo-go-driver

```
  "mongo": {
    "dsn": "mongodb://root:123456@127.0.0.1:27017,127.0.0.1:27017/?authSource=admin",
    "minPoolSize": 10,
    "maxPoolSize": 100
  }
## DSN
## maxIdleConn：db连接池最大空闲链接数
## maxOpenConn：db连接池最大连接数
## connMaxLifetime：连接池超时时间(单位: time.Hour)
```

##

#### kafka brokers

```
  "kafka": {
    "brokers": [
      "192.168.3.40:9092",
      "192.168.3.41:9092",
      "192.168.3.42:9092"
    ],
    "cert_file": "",
    "key_file": "",
    "ca_file": "",
    "verify_ssl": false
  }
## brokers：节点
## cert_file：证书文件
## key_file：证书文件
## ca_file：证书文件
## verify_ssl：是否开启ssl验证
```

##

#### zipkin config

```
  "zipkin": {
    "service_name": "laya-go",
    "service_endpoint": "0.0.0.0:10080",
    "zipkin_addr": "http://zipkin.xxx.cn/api/v2/spans",
    "mod": 1
  }
## service_name：服务名称
## service_endpoint = "0.0.0.0:10080"
## zipkin_addr = "http://zipkin.xxx.cn/api/v2/spans"
## mod = 1 //1==全量 值越大，采样率月底，对性能影响越小
```
