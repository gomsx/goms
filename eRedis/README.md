
# eRedis

Cache, 使用 Redis.


## 运行服务
```
cd goms/eRedis/cmd

# 使用默认配置文件
go run . &  

# 使用指定配置文件
go run . & -confpath=../configs  
```

## 测试API

http
```
# 使用 http 方法 GET /ping
curl localhost:8080/ping

# 使用 http 方法 GET /ping, 参数 message=xxx
curl localhost:8080/ping?message=xxx

# 使用 http 方法 POST /user/user, 参数 name=xxx sex=0
curl -X POST -d "uid=123&name=xxx&sex=1" localhost:8080/user

# 使用 http 方法 PUT /user/user, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=xxx&sex=1" localhost:8080/user/123

# 使用 http 方法 GET /user/user, 参数 uid=123
curl -X GET localhost:8080/user/123
curl -X GET localhost:8080/user?uid=123

# 使用 http 方法 DELETE /user/user, 参数 uid=123
curl -X DELETE localhost:8080/user/123
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 service.goms.User/Ping
```