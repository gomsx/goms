
# eRedis

这个模块加入缓存, 使用 Redis.


## 生成代码

protoc
```
cd goms/eRedis/api/pb

# 执行: pb.go 文件头的指令
go generate ./pb.go 
```
>pb.go 文件   
//go:generate protoc --go_out=plugins=grpc:../ call.proto
package pb


mockgen
```
cd goms/eRedis/api/mock

# 执行 mock.go 文件头的指令 
go generate ./mock.go
```
>mock.go 文件  
//go:generate mockgen  -package mock -destination ./callclient_mock.go  github.com/fuwensun/goms/eRedis/api CallClient


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
# 使用 http 方法 GET /call/ping
curl  localhost:8080/call/ping

# 使用 http 方法 GET /call/ping, 参数 message=xxx
curl  localhost:8080/call/ping?message=xxx

# 使用 http 方法 GET /user/user, 参数 \?uid\=12
curl  localhost:8080/user/readname\?uid\=12

# 使用 http 方法 POST /user/user, 参数 name=xxx sex=0
curl  -X POST localhost:8080/user/user\?&name\=xxx\&sex\=0

# 使用 http 方法 PUT /user/user, 参数 uid=123 name=yyy sex=1
curl  -X PUT localhost:8080/user/user\?uid\=123\&name\=yyy\&sex\=1

# 使用 http 方法 GET /user/user, 参数 uid=123
curl  -X GET localhost:8080/user/user\?uid\=123

# 使用 http 方法 DELETE /user/user, 参数 uid=123
curl  -X DELETE localhost:8080/user/user\?uid\=123
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 api.Call/Ping, 带参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Call/Ping 

```