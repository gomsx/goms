# eHttp

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| ✖ | 缓存
 ✖ | grpc服务| ✖ | 日志
 ✖ | 读取配置| ✖ | 测试
 ✖ | 数据库  | ✖ | API管理

## 运行服务

```
cd goms/eHttp/cmd

go run . & 
```

## 测试 API

http
```
# 使用 http 方法 GET /ping
curl localhost:8080/ping

# 使用 http 方法 GET /ping, 参数 message=xxx
curl localhost:8080/ping?message=xxx
```
