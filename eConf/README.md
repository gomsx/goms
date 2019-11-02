

## 生成代码

protoc
```
cd goms/eConf/api/pb

go generate ./pb.go 
# 实际执行: protoc --go_out=plugins=grpc:../ ./call.proto
```

mockgen
```
cd goms/eConf/api/mock

go generate ./mock.go
# 实际执行: mockgen  -package mock -destination ./callclient_mock.go \
github.com/fuwensun/goms/eConf/api CallClient
```


## 运行服务
```
cd goms/eConf/cmd

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
```
grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 api.Call/Ping
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Call/Ping 

```