运行服务
```
cd goms/eGrpc/cmd

go run . & 

```

测试API
```
curl  localhost:8080/test/ping              # 使用 http 方法 /test/ping

grpcurl -plaintext localhost:50051 list     # 获取 grpc 方法列表

grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Call/Ping   # 使用 grpc 方法 api.Call/Ping

```