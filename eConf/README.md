# eConf

配置,从本地配置文件读取.

## 生成代码

protoc
```
cd goms/eConf/api/pb

# 执行 pb.go 文件头的指令
go generate ./pb.go 
```
>pb.go 文件   
//go:generate protoc --go_out=plugins=grpc:../ call.proto
package pb


mockgen
```
cd goms/eConf/api/mock

# 执行 mock.go 文件头的指令 
go generate ./mock.go
```
>mock.go 文件  
//go:generate mockgen  -package mock -destination ./callclient_mock.go  github.com/fuwensun/goms/eConf/api CallClient


## 运行服务
```
cd goms/eConf/cmd

# 使用默认的配置文件
go run . &  

# 使用指定的配置文件
go run . & -confpath=../configs  
```


## 测试API

http
```
# 使用 http 方法 GET　/call/ping
curl  localhost:8080/call/ping

# 使用 http 方法 GET /call/ping, 参数 message=xxx
curl  localhost:8080/call/ping?message=xxx
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 api.Call/Ping, 参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Call/Ping 

```