# eRedis

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| ✔ | 缓存
 ✔ | grpc服务| &nbsp; | 测试
 ✔ | 数据库  | &nbsp; | API管理

## 概念

### 缓存策略

- Cache-Aside
- Read-Though Cache
- Write-Through Cache
- Write-Around
- Write-Back

### 缓存过期策略

- FIFO (First In First out): 先进先出,淘汰最早进来的
- LRU (Least recently used): 最早使用的,淘汰最早使用的
- LFU (Least frequently used): 最不常使用的,淘汰使用次数最少的

## 成果

### 运行服务

```
cd goms/eRedis/cmd

# 使用默认配置文件
go run . &  

# 使用指定配置文件
go run . & -cfgpath=../configs  
```

### 测试(使用) API

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
grpcurl -plaintext -d '{"message":"xxx"}'  localhost:50051 service.goms.User/Ping

# 使用 grpc 方法 service.goms.User/CreateUser, 参数 {"name":"xxx","sex":"1"}
grpcurl -plaintext -d '{"name":"xxx","sex":"1"}' localhost:50051 service.goms.User/CreateUser

# 使用 grpc 方法 service.goms.User/ReadUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"uid":"123"}' localhost:50051 service.goms.User/ReadUser

# 使用 grpc 方法 service.goms.User/UpdateUser, 参数 {"uid":"123","name":"xxx","sex":"1"} 
grpcurl -plaintext -d '{"uid":"123","name":"xxx","sex":"1"}' localhost:50051 service.goms.User/UpdateUser

# 使用 grpc 方法 service.goms.User/DeleteUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"uid":"123"}' localhost:50051 service.goms.User/DeleteUser
```

## 参考

https://www.jianshu.com/p/207130233e60  
https://www.cnblogs.com/s-b-b/p/6047954.html  
