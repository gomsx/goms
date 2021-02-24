# eHttp

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| &nbsp; | 缓存
 &nbsp; | grpc服务| &nbsp; | 测试
 &nbsp; | 数据库  | &nbsp; | API管理

## 概念

## 成果
### 运行服务

```
cd goms/eHttp/cmd

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
