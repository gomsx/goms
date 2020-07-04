# eGrpc

rpc 服务,使用 grpc 包.
## 生成代码

protoc
```
cd goms/eGrpc/api/grpc/pb

# 执行 pb.go 文件头的指令
go generate ./pb.go 
```
>pb.go 文件   
//go:generate protoc --go_out=plugins=grpc:../ api.proto

## 运行服务
```
cd goms/eGrpc/cmd

go run . & 
```

## 测试API

http
```
# 使用 http 方法 GET /ping
curl  localhost:8080/ping

# 使用 http 方法 GET /ping, 参数 message=xxx
curl  localhost:8080/ping?message=xxx
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 service.goms.User/Ping 
```

