# eHttp

http 服务,使用 gin 框架.

## 运行服务
```
cd goms/eHttp/cmd

go run . & 

```

## 测试API
http
```
# 使用 http 方法 GET /call/ping
curl  localhost:8080/call/ping

# 使用 http 方法 GET /call/ping, 参数 message=xxx
curl  localhost:8080/call/ping?message=xxx
```