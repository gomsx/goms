
# eRedis

这个模块加入缓存, 使用 Redis.


## 生成代码

protoc
```
cd goms/eRedis/api/pb

# 实际执行: pb.go 文件头注释部分指令 
go generate ./pb.go 
```
pb.go 文件

> //go:generate protoc --go_out=plugins=grpc:../ call.proto
package pb


mockgen
```
cd goms/eRedis/api/mock

# 实际执行: mock.go 文件头注释部分指令 
go generate ./mock.go
```
mock.go 文件
> //go:generate mockgen  -package mock -destination ./callclient_mock.go  github.com/fuwensun/goms/eRedis/api CallClient
package moc


## 运行服务
```
cd goms/eRedis/cmd

# 使用默认的配置文件路径
go run . &  

# 使用指定的配置文件路径
go run . & -confpath=../configs  
```


## 测试API

http
```
# 使用 http 方法 /call/ping
curl  localhost:8080/call/ping

# 使用 http 方法 /call/ping, 带参数 message=xxx
curl  localhost:8080/call/ping?message=xxx

# 使用 http 方法 /user/readname, 带参数 \?uid\=12
curl  localhost:8080/user/readname\?uid\=12

# 使用 http 方法 /user/updatename, 带参数 \?uid\=12\&name\=jieke
curl  localhost:8080/user/updatename\?uid\=12\&name\=jieke
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 api.Call/Ping, 带参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Call/Ping 

```