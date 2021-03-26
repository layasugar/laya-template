###配置文件说明

#### 项目整体配置
```
## 项目名称
app_name = "laya_go_template"
## 监听端口
http_listen = "0.0.0.0:10080"
## mode debug or release
app_mode = "debug"
## 当前项目请求url
app_url = "https://github.com/layatips/laya-go-template"
## gin的http请求日志输出到文件
gin_log = "/home/logs/app/laya-go-template/gin_http.log"
## 是否打印入参和出参
params_log = true
```

##
### log
#### https://github.com/uber-go/zap
```
## log driver
driver = "file"
## path
path = "./logs/log.info"
## MB
max_size = 32
## 保留旧文件的最大天数
max_age = 90
## 保留旧文件的最大个数
max_backups = 300
```

##
### cache
#### https://github.com/patrickmn/go-cache
```
open = true
##秒-默认过期时间
default_exp=300
##秒-每多少清除一次过期物品
cleanup=600
```

##
### mysql
#### http://gorm.io/zh_CN/docs/index.html
```
## open
open = true
## 驱动类型,gorm支持的驱动类型有MySQL, PostgreSQL, SQlite, SQL Server
driver = "mysql"
## db DSN
dsn = "root:123456@tcp(127.0.0.1:3306)/lata_test?charset=utf8&parseTime=True&loc=Local"
## db连接池最大空闲链接数
maxIdleConn = 10
## db连接池最大连接数
maxOpenConn = 100
## 连接池超时时间(单位: time.Hour)
connMaxLifetime = 6
```

##
### redis
#### https://github.com/go-redis/redis
```
## open
open = true
## db
db = 0
## addr-链接DSN
addr = "127.0.0.1:6379"
## pwd-链接密码
pwd = ""
## poolSize-链接池大小
poolSize = 5
## maxRetries-链接最大重试次数
maxRetries = 3
## idleTimeout-空闲链接超时时间(单位：time.Second秒)
idleTimeout = 1000
```

## 
### mongo
#### https://github.com/mongodb/mongo-go-driver
```
## open
open = true
## DSN
dsn = "mongodb://root:123456@127.0.0.1:27017,127.0.0.1:27017/?authSource=admin"
## db连接池最大空闲链接数
maxIdleConn = 10
## db连接池最大连接数
maxOpenConn = 100
## 连接池超时时间(单位: time.Hour)
connMaxLifetime = 6
```

##
#### kafka brokers
```
open = true
brokers = ["192.168.3.40:9092", "192.168.3.41:9092"]
cert_file = ""
key_file = ""
ca_file = ""
verify_ssl = false
```

##
#### zipkin config
```
open = true
service_name = "payment"
service_endpoint = "0.0.0.0:10080"
zipkin_addr = "http://zipkin.xthklocal.cn/api/v2/spans"
mod = 1 //1==全量。值越大，采样率月底，对性能影响越小
```
