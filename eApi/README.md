# eApi

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| ✔ | 缓存
 ✔ | grpc服务| ✔ | 日志
 ✔ | 读取配置| ✔ | 测试
 ✔ | 数据库  | ✔ | API管理

## 设计原则

- 使用成熟度合适的 RESTful API
- 避免简单封装
- 关注点分离
- 完全穷尽,彼此独立(正交化？)
- 版本化
- 合理命名
- 安全

>https://zhuanlan.zhihu.com/p/86446096  
https://docs.microsoft.com/zh-cn/azure/architecture/best-practices/api-design  
## 类型

- http
- grpc

## http api

- RESTful 风格

- [OpenAPI 规范][23]

- [Swagger 工具][24]

- 设计方法
  - 文档优先，先文档后代码,工具 go-swagger
  - 代码优先，先代码后文档,工具 swaggo

[23]:https://github.com/OAI/OpenAPI-Specification  
[24]:https://swagger.io/  

## grpc api

- [protocol buffers 协议][31]
  - [Protocol Compiler][32]
  - [Protobuf Runtime][33]

- [grpc 框架][41]
  - go 实现 [grpc-go][42]
  - grpc 组件 [grpc-ecosystem][43]

[31]:https://developers.google.com/protocol-buffers
[32]:https://github.com/protocolbuffers/protobuf
[33]:https://github.com/protocolbuffers/protobuf-go
[41]:https://www.grpc.io
[42]:https://github.com/grpc/grpc-go
[43]:https://github.com/grpc-ecosystem

## 运行服务

```
cd goms/eApi/cmd

# 使用默认配置文件
go run . &  

# 使用指定配置文件
go run . & -cfgpath=../configs  
```

## 测试API
log
```
curl localhost:8080/v1/logs/all

curl localhost:8080/v1/logs?name=all

curl -X PUT -d "level=info" localhost:8080/v1/logs/all

curl -X PUT -d "name=all&level=info" localhost:8080/v1/logs
```

http
```
# 使用 http 方法 GET /v1/ping
curl localhost:8080/v1/ping

# 使用 http 方法 GET /v1/ping, 参数 message=xxx
curl localhost:8080/v1/ping?message=xxx

# 使用 http 方法 POST /v1/users, 参数 name=xxx sex=1
curl -X POST -d "name=xxx&sex=1" localhost:8080/v1/users

# 使用 http 方法 GET /v1/users, 参数 uid=123
curl -X GET localhost:8080/v1/users/123
curl -X GET localhost:8080/v1/users?uid=123

# 使用 http 方法 PUT /v1/users, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=yyy&sex=1" localhost:8080/v1/users/123

# 使用 http 方法 DELETE /v1/users, 参数 uid=123
curl -X DELETE localhost:8080/v1/users/123
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.v1.User/Ping, 参数 {"message":"xxx"}
grpcurl -plaintext -d '{"data":{"message":"xxx"}}' localhost:50051 service.goms.v1.User/Ping

# 使用 grpc 方法 service.goms.v1.User/CreateUser, 参数 {"name":"xxx","sex":"1"}
grpcurl -plaintext -d '{"data":{"name":"xxx","sex":"1"}}' localhost:50051 service.goms.v1.User/CreateUser

# 使用 grpc 方法 service.goms.v1.User/ReadUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"data":{"uid":"123"}}' localhost:50051 service.goms.v1.User/ReadUser

# 使用 grpc 方法 service.goms.v1.User/UpdateUser, 参数 {"uid":"123","name":"yyy","sex":"1"}
grpcurl -plaintext -d '{"data":{"name":"yyy","sex":"1","uid":"123"}}' localhost:50051 service.goms.v1.User/UpdateUser

# 使用 grpc 方法 service.goms.v1.User/DeleteUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"data":{"name":"xxx","sex":"1","uid":"123"}}' localhost:50051 service.goms.v1.User/DeleteUser
```

gateway
```
curl -X GET localhost:8081/v1/ping/

curl -X GET localhost:8081/v1/ping/xxx

curl -X POST -d '{"data":{"name":"xxx","sex":"1"}}' localhost:8081/v1/users

curl -X GET localhost:8081/v1/users/123

curl -X PUT -d '{"data":{"name":"yyy","sex":"1","uid":"123"}}' localhost:8081/v1/users

curl -X DELETE localhost:8081/v1/users/123
```
