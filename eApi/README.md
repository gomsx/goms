# eApi

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| ✔ | 缓存
 ✔ | grpc服务| ✔ | 测试
 ✔ | 数据库  | ✔ | API管理

## 概念

### 设计原则

- 使用成熟度合适的 RESTful API
- 避免简单封装
- 关注点分离
- 完全穷尽,彼此独立(正交化？)
- 版本化
- 合理命名
- 安全


### 类型

- http
- grpc

### HTTP API

- RESTful 风格
- [OpenAPI 规范][23]
- [Swagger 工具][24]
- 设计方法
  - 文档优先，先文档后代码,工具 go-swagger
  - 代码优先，先代码后文档,工具 swaggo

[23]:https://github.com/OAI/OpenAPI-Specification  
[24]:https://swagger.io/  

### GRPC API

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

## 依赖

### grpc-gateway

```
# 安装
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2

# 查看
ls $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
protoc-gen-grpc-gateway
protoc-gen-openapiv2

ls $GOPATH/bin
protoc-gen-grpc-gateway
protoc-gen-openapiv2

# 使用
protoc --grpc-gateway_out=logtostderr=true:. *.proto
protoc --swagger_out=logtostderr=true:. *.proto
```

### go-swagger  

```
# 安装
go get -u github.com/go-swagger/go-swagger/cmd/swagger

# 查看
ls $GOPATH/bin
swagger

# 使用
swagger serve --host=0.0.0.0 --port=9000 --no-open api.swagger.json

# 访问
http://localhost:9000/docs
```

## 成果

### 运行服务

```
cd goms/eApi/cmd

# 使用默认配置文件
go run . &  

# 使用指定配置文件
go run . & -cfgpath=../configs  
```

### 测试(使用) API

log
```
curl localhost:8080/v1/log

curl -X PUT -d "level=info" localhost:8080/v1/log
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

## 参考

https://zhuanlan.zhihu.com/p/86446096  
https://docs.microsoft.com/zh-cn/azure/architecture/best-practices/api-design  
