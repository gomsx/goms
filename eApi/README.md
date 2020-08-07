# eApi

### grpc-gateway

```
# 安装
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# 查看
ls $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
protoc-gen-grpc-gateway
protoc-gen-swagger

ls $GOPATH/bin
protoc-gen-grpc-gateway
protoc-gen-swagger

# 使用
protoc --grpc-gateway_out=logtostderr=true:. *.proto
protoc --swagger_out=logtostderr=true:. *.proto
```

>https://grpc-ecosystem.github.io/grpc-gateway/  
https://github.com/grpc-ecosystem/grpc-gateway　　

### swagger  

```
# 安装
go get -u github.com/go-swagger/go-swagger/cmd/swagger

# 查看
ls $GOPATH/bin
swagger

# 使用
swagger serve --host=0.0.0.0 --port=9000 --no-open api.swagger.json

# 访问
http://localhost:9000/docs
```

>https://github.com/go-swagger/go-swagger  

## 测试API
log
```
curl localhost:8080/v1/logs/all
curl localhost:8080/v1/logs?name=all
curl -X PUT -d "level=info" localhost:8080/v1/logs/all
curl -X PUT -d "name=all&level=info" localhost:8080/v1/logs
```

http
```
# 使用 http 方法 GET /v1/ping
curl localhost:8080/v1/ping 

# 使用 http 方法 GET /v1/ping, 参数 message=xxx
curl localhost:8080/v1/ping?message=xxx 

# 使用 http 方法 POST /v1/users, 参数 name=xxx sex=1
curl -X POST -d "name=xxx&sex=1" localhost:8080/v1/users 

# 使用 http 方法 GET /v1/users, 参数 uid=123
curl -X GET localhost:8080/v1/users/123 
curl -X GET localhost:8080/v1/users?uid=123 

# 使用 http 方法 PUT /v1/users, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=xxx&sex=1" localhost:8080/v1/users/123 

# 使用 http 方法 DELETE /v1/users, 参数 uid=123
curl -X DELETE localhost:8080/v1/users/123 
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.v1.User/Ping, 参数 {"message":"xxx"}
grpcurl -plaintext -d '{"message":"xxx"}'  localhost:50051 service.goms.v1.User/Ping

# 使用 grpc 方法 service.goms.v1.User/CreateUser, 参数 {"name":"xxx","sex":"1"}
grpcurl -plaintext -d '{"name":"xxx","sex":"1"}' localhost:50051 service.goms.v1.User/CreateUser

# 使用 grpc 方法 service.goms.v1.User/ReadUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"uid":"123"}' localhost:50051 service.goms.v1.User/ReadUser

# 使用 grpc 方法 service.goms.v1.User/UpdateUser, 参数 {"uid":"123","name":"xxx","sex":"1"} 
grpcurl -plaintext -d '{"uid":"123","name":"xxx","sex":"1"}' localhost:50051 service.goms.v1.User/UpdateUser

# 使用 grpc 方法 service.goms.v1.User/DeleteUser, 参数 {"uid":"123"}
grpcurl -plaintext -d '{"uid":"123"}' localhost:50051 service.goms.v1.User/DeleteUser
```

gateway
```
curl -X GET localhost:8081/v1/ping

curl -X GET localhost:8081/v1/ping?message=xxx

curl -X POST -d '{"name":"xxx","sex":1}' localhost:8081/v1/users

curl -X GET localhost:8081/v1/users/123456

curl -X PUT -d '{"name":"xxx","sex":1,"uid":123456}' localhost:8081/v1/users

curl -X DELETE localhost:8081/v1/users/123456 
```

