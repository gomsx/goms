# eTest

go 单元测试
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
curl localhost:8080/ping -w "\n"

# 使用 http 方法 GET /ping, 参数 message=xxx
curl localhost:8080/ping?message=xxx -w "\n"

# 使用 http 方法 POST /users, 参数 name=xxx sex=0
curl -X POST -d "name=xxx&sex=1" localhost:8080/users -w "\n"

# 使用 http 方法 PUT /users, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=xxx&sex=1" localhost:8080/users/123 -w "\n"

# 使用 http 方法 GET /users, 参数 uid=123
curl -X GET localhost:8080/users/123 -w "\n"
curl -X GET localhost:8080/users?uid=123 -w "\n"

# 使用 http 方法 DELETE /users, 参数 uid=123
curl -X DELETE localhost:8080/users/123 -w "\n" 
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 service.goms.User/Ping

# 使用 grpc 方法 service.goms.User/CreateUser, 参数 {"Name": "xxx","Sex":"0"}
grpcurl -plaintext -d '{"Name": "xxx","Sex":"0"}' localhost:50051 service.goms.User/CreateUser

# 使用 grpc 方法 service.goms.User/UpdateUser, 参数 {"Uid":"123","xxx":"name","Sex":"1"} 
grpcurl -plaintext -d '{"Uid":"123","xxx":"name","Sex":"1"}' localhost:50051 service.goms.User/UpdateUser

# 使用 grpc 方法 service.goms.User/ReadUser, 参数 {"Uid":"123"}
grpcurl -plaintext -d '{"Uid":"123"}' localhost:50051 service.goms.User/ReadUser

# 使用 grpc 方法 service.goms.User/DeleteUser, 参数 {"Uid":"123"}
grpcurl -plaintext -d '{"Uid":"123"}' localhost:50051 service.goms.User/DeleteUser
```