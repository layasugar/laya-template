### laya-go
gin+gorm+go-redis为基本骨架打造的目录结构，方便开发，开箱即用

#### web应用

- 参考mian.go

#### 服务应用

```
// ServerSetup 初始化服务设置
func ServerSetup() *xgs.App {
	app := xgs.NewApp(xgs.SetLogger)
	//app.Use(conf.Load)
	return app
}

func main() {
	app := ServerSetup()
	
    // your code

	app.Use(listenHttp)
}

func listenHttp() {
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(utils.RespSuc))
	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(utils.RespSuc))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(utils.RespSuc))
	})
	log.Printf("http_listen: %s", genv.HttpListen())
	_ = http.ListenAndServe(genv.HttpListen(), nil)
}
```

#### 配置文件的使用
- 运行服务时配置，具体配置请参照conf/app.json，按照结构配置，base和log和mem提供默认配置可以不配
- 别问为什么用json，问就是不知道
- 参考[配置说明](https://github.com/layatips/laya-go/tree/master/conf)

#### api模板

- [github链接](https://github.com/layatips/laya-go)
- 总的来说就是自己的日常开发形成的一套习惯，代码拉下来就可以运行，方便新api或者web项目的快速搭建 组件都是目前golang比较火的组件

#### 启动运行

1. clone代码
    - ```git clone https://github.com/layatips/laya-go.git```
2. 修改配置conf/app.json
3. 启动
    - ```go run github.com/layatips/laya-go```
4. docker方式启动
    - ```docker build -t laya-go:1.0 . && docker run --name laya-go -p 10080:10080 --network devops --network-alias laya-go laya-go:1.0 ```

##   

#### route

- https://github.com/gin-gonic/gin
- gin的原来的方式创建路由没有任何修改
- 参考routes/base.go

##

#### controller

- 参考controller/base.go

##

#### model

- 传统的ddd模式
- page业务层
- data数据层
- dao连接或者请求层

##

#### middleware

- 请参考gin的middleware
- ```app.WebServer.Use(middleware.LogInParams)```

##

#### log

- uber开源的zap日志[github链接](https://github.com/uber-go/zap)
- 使用方式一（会打印request_id）

```
glogs.InfoFR(c,format,...interface)
glogs.WarnFR(c,format,...interface)
glogs.ErrorFR(c,format,...interface)
```

- 使用方式二（不会打印request_id）

```
glogs.InfoF(format,...interface)
glogs.WarnF(format,...interface)
```

- 可自己拓展

#### zipkin链路追踪

- 开源的zipkin的sdk[github链接](github.com/openzipkin/zipkin-go)
- 使用

```
span := glogs.StartTrace(c,funcName)
your code ...
glogs.StopTrace(span)
```

#### 钉钉通知

- 使用

```
var d = &glogs.AlarmData{
  RobotKey: robotKey,
  RobotHost: robotHost,
  Title:       "这是标题",
  Description: "这是问题描述",
  Content: map[string]interface{}{
      "adasdasd": "sadasssss",
      "sadddd":   "11111111",
  }}
glogs.SendAlarm(d)
```

##

#### cache(go实现的memcache)

- [github链接](https://github.com/patrickmn/go-cache)
- 使用```data.GetMem(),data.SetMem(),data.DelMem()```

##

#### 错误代码

- 定义在global/errno/里面
- 使用utils.SystemErr
- 返回的是error类型，response里面会根据error不同而加载错误信息

##

#### utils

- time.go是基于sql.nulltime自己实现的
- helps是辅助工具类
- redis_lock是实现分布式的redis锁

##

#### 协程池可参考

- ants.pool

```
p, _ := ants.NewPoolWithFunc(runtime.NumCPU()*2, func(data interface{}) {
   if pData, ok := data.([]OtsOrder); ok {
       err := batchWriteOtsOrder(pData)
       if err != nil {
           glogs.Error("WorkerRunOnlineCourse err:%v", err.Error())
       }
   }
})
defer p.Release()
// Submit tasks one by one.
for data := range BatchOtsOrderChan {
   _ = p.Invoke(data)
}
```

#### 流批处理可参考
```
func WorkerReadyClass(inChan <-chan PutData, batchSize int, duration time.Duration) {
	batchData := make([]PutData, 0, batchSize)
	i := 0
	for {
		ctx, _ := context.WithTimeout(context.Background(), duration)
		select {
		case data, ok := <-inChan:
			if !ok {
				return
			}
			batchData = append(batchData, data)
			i++
			if i >= batchSize {
				WorkChanClass <- batchData
				batchData = make([]PutData, 0, batchSize)
				i = 0
			}
		case <-ctx.Done():
			if len(batchData) > 0 {
				WorkChanClass <- batchData
				batchData = make([]PutData, 0, batchSize)
				i = 0
			}
		}
	}
}
```

#### 推荐工具

##### 数据库直接生成gorm的struct

- [github链接](https://github.com/Shelnutt2/db2struct)
- ```db2struct --host localhost -d database --package db -p 123456 --user root --guregu --gorm -t tableName --struct structName```

##### hey 压测工具

- [github链接](https://github.com/rakyll/hey)
- ```hey -n 100 -c 1000```