# eGrpc

rpc 服务,使用 grpc 包.

## 运行服务
```
cd goms/eGrpc/cmd

go run . & 

```

## 测试API

http
```
# 使用 http 方法 /call/ping
curl  localhost:8080/call/ping

# 使用 http 方法 /call/ping, 带参数 message=xxx
curl  localhost:8080/call/ping?message=xxx
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 api.Call/Ping, 带参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Call/Ping 
``