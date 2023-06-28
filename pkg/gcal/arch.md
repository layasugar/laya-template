### 四个对象
- service: 服务发现。定期更新服务提供方的信息（Host List）
- convert: 请求数据的序列化和返回数据的格式化 (form, json, mcpack2, raw)
	- Pack(interface{}, *context.Context) ([]byte, error)
	- UnPack(interface{}, interface{}, *context.Context) error
- protocol: 请求发送的对象 (http, nshead, pbrpc)
 	- CreateProtocolContext(service.Service, interface{}, *context.Context) (interface{}, error)
	- DoRequest(service.Service, interface{}, *context.Context) (Response, error)
- balance： 负载均衡策略
	- FetchServer(s service.Service) (*service.Addr, error)


### 基本流程

```
func Do(serviceName string, request interface{}, response interface{}, converterType ConverterType) error
```

- 根据 serviceName 得到 service 对象 S
- 根据 S 的 protocol 设定找到 protocol P
- P CreateProtocolContext 
    - 得到 Converter 对象 C 进行 Pack
    - 根据 balance 策略得到 具体的host
- P DoRequest 得到数据
- C 进行 UnPack

