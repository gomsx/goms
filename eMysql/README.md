

## 生成代码

protoc
```
cd goms/eMysql/api/pb

# 实际执行: protoc --go_out=plugins=grpc:../ ./test.proto
go generate ./pb.go 
```

mockgen
```
cd goms/eMysql/api/mock

# 实际执行: mockgen  -package mock -destination ./testclient_mock.go \
github.com/fuwensun/goms/eMysql/api TestClient
go generate ./mock.go
```


## 运行服务
```
cd goms/eMysql/cmd

# 使用默认的配置文件路径
go run . &  

# 使用指定的配置文件路径
go run . & -confpath=../configs  
```


## 测试API

http
```
# 使用 http 方法 /test/ping
curl  localhost:8080/test/ping

# 使用 http 方法 /test/ping, 带参数 message=xxx
curl  localhost:8080/test/ping?message=xxx
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 api.Test/Ping, 带参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Test/Ping 

```