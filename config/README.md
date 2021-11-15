### 配置文件说明

### 注意：请使用一个配置文件

#### 项目app基础配置

```json
{
  "app": {
    "name": "laya-go",
    "mode": "dev1",
    "run_mode": "debug",
    "http_listen": "0.0.0.0:80",
    "url": "http://127.0.0.1:80",
    "pprof": true,
    "params": true,
    "logger": "/home/logs/app",
    "version": "1.0.0"
  }
}
```

- app.name：项目名称
- app.mode：应用环境
- app.run_mode: debug or release(gin的运行模式)
- app.http_listen: gin程序监听的端口
- app.url: 程序外网路径
- app.pprof: 是否开启pprof分析
- app.params: 是否开启入参出参打印
- app.logger: 日志路径
- app.version: 程序版本
