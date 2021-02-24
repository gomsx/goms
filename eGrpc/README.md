# eGrpc

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| &nbsp; | 缓存
 ✔ | grpc服务| &nbsp; | 测试
 &nbsp; | 数据库  | &nbsp; | API管理

## 概念

### protobuffer

### grpc

## 依赖

### 生成代码

protoc
```
cd goms/eGrpc/api/grpc/pb

# 执行 pb.go 文件头的指令
go generate ./pb.go 
```

## 成果

### 运行服务

```
cd goms/eGrpc/cmd

go run . & 
```

### 测试(使用) API

http
```
# 使用 http 方法 GET /ping
curl localhost:8080/ping

# 使用 http 方法 GET /ping, 参数 message=xxx
curl localhost:8080/ping?message=xxx
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"message":"xxx"}
grpcurl -plaintext -d '{"message":"xxx"}' localhost:50051 service.goms.User/Ping 
```
