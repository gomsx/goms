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
cd goms/eRedis/cmd

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

# 使用 http 方法 POST /user/user, 参数 name=xxx sex=0
curl -X POST -d "name=xxx&sex=1" localhost:8080/user -w "\n"

# 使用 http 方法 PUT /user/user, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=xxx&sex=1" localhost:8080/user/123 -w "\n"

# 使用 http 方法 GET /user/user, 参数 uid=123
curl -X GET localhost:8080/user/123 -w "\n"
curl -X GET localhost:8080/user?uid=123 -w "\n"

# 使用 http 方法 DELETE /user/user, 参数 uid=123
curl -X DELETE localhost:8080/user/123 -w "\n" 
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 service.goms.User/Ping
```