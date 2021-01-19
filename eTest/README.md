# eTest

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| ✔ | 缓存
 ✔ | grpc服务| ✔ | 日志
 ✔ | 读取配置| ✔ | 测试
 ✔ | 数据库  | ✖ | API管理


## 依赖

单元测试依赖通过伪造模式解决，经典的伪造模式有:

- 桩对象 stub
- 模拟对象 mock
- 伪对象 fake

## 框架

- 轻量级的 testing 包
- 更多功能的 [goconvey][21] 框架

Stub, Mock, Fakes等工具来隔离用例和依赖.

- 全局变量通过 [GoStub][22] 框架打桩
- 接口通过 [GoMock][23] 框架打桩
- 方法/函数通过 [Monkey][24] 框架打桩

[21]:https://github.com/smartystreets/goconvey
[22]:https://github.com/prashantv/gostub
[23]:https://github.com/golang/mock
[24]:https://github.com/bouk/monkey

## 运行服务

```
cd goms/eTest/cmd

# 使用默认配置文件
go run . &  

# 使用指定配置文件
go run . & -cfgpath=../configs  
```

## 测试API

http
```
# 使用 http 方法 GET /ping
curl localhost:8080/v1/ping 

# 使用 http 方法 GET /ping, 参数 message=xxx
curl localhost:8080/v1/ping?message=xxx 

# 使用 http 方法 POST /users, 参数 name=xxx sex=1
curl -X POST -d "name=xxx&sex=1" localhost:8080/v1/users 

# 使用 http 方法 GET /users, 参数 uid=123
curl -X GET localhost:8080/v1/users/123 
curl -X GET localhost:8080/v1/users?uid=123 

# 使用 http 方法 PUT /users, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=xxx&sex=1" localhost:8080/v1/users/123 

# 使用 http 方法 DELETE /users, 参数 uid=123
curl -X DELETE localhost:8080/v1/users/123 
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"message":"xxx"}
grpcurl -plaintext -d '{"message":"xxx"}' localhost:50051 service.goms.User/Ping

# 使用 grpc 方法 service.goms.User/CreateUser, 参数 {"name":"xxx","sex":"1"}
grpcurl -plaintext -d '{"name":"xxx","sex":"1"}' localhost:50051 service.goms.User/CreateUser

# 使用 grpc 方法 service.goms.User/ReadUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"uid":"123"}' localhost:50051 service.goms.User/ReadUser

# 使用 grpc 方法 service.goms.User/UpdateUser, 参数 {"uid":"123","name":"xxx","sex":"1"} 
grpcurl -plaintext -d '{"uid":"123","name":"xxx","sex":"1"}' localhost:50051 service.goms.User/UpdateUser

# 使用 grpc 方法 service.goms.User/DeleteUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"uid":"123"}' localhost:50051 service.goms.User/DeleteUser
```
