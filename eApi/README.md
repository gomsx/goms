# eApi




## 安装


### Protocol Compiler Installation

https://github.com/protocolbuffers/protobuf  
https://github.com/protocolbuffers/protobuf/releases  
https://developers.google.com/protocol-buffers/docs/overview  
```
sudo apt install protobuf-compiler
protoc --version
```
>下载 Protocol 编译器 protoc

### Protobuf Runtime Installation
Go support for Protocol Buffers

https://github.com/golang/protobuf  
https://developers.google.com/protocol-buffers/docs/gotutorial  

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```
>下载 protobuf 的 go 运行时，包含 protoc 编译器的插件和与之配对的包，如 protoc-gen-go 插件和 github.com/golang/protobuf/proto 包，还有 protoc-gen-go/grpc 插件和 google.golang.org/grpc 包(不在这个运行时中).并安装 protoc-gen-go 插件到 GOPATH 目录下,命令"protoc --go_out=:. *.proto"中的"--go_out"参数可用，包含静态子插件 grpc，命令"protoc --go_out=plugins=grpc:. *.proto"中的"plugins=grpc"参数可用

### grpc Package Installation

https://github.com/grpc/grpc-go  
https://www.grpc.io/docs/tutorials/basic/go/  

```
go get -u google.golang.org/grpc
```
>这个包是 grpc 的 go 实现，已经映射为 github.com/grpc/grpc-go.

###　grpc-gateway

https://grpc-ecosystem.github.io/grpc-gateway/　　
https://github.com/grpc-ecosystem/grpc-gateway　　

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```


## 测试API

http
```
# 使用 http 方法 GET /ping
curl localhost:8080/ping -w "\n"

# 使用 http 方法 GET /ping, 参数 message=xxx
curl localhost:8080/ping?message=xxx -w "\n"

# 使用 http 方法 POST /user/user, 参数 name=xxx sex=0
curl -X POST -d "name=xxx&sex=1" localhost:8080/user -w "\n"

# 使用 http 方法 PUT /user/user, 参数 uid=123 name=yyy sex=1
curl -X PUT -d "name=xxx&sex=1" localhost:8080/user/123 -w "\n"

# 使用 http 方法 GET /user/user, 参数 uid=123
curl -X GET localhost:8080/user/123 -w "\n"
curl -X GET localhost:8080/user?uid=123 -w "\n"

# 使用 http 方法 DELETE /user/user, 参数 uid=123
curl -X DELETE localhost:8080/user/123 -w "\n" 
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"Message": "xxx"}
grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 service.goms.User/Ping

# 使用 grpc 方法 service.goms.User/CreateUser, 参数 {"Name": "xxx","Sex":"0"}
grpcurl -plaintext -d '{"Name": "xxx","Sex":"0"}' localhost:50051 service.goms.User/CreateUser

# 使用 grpc 方法 service.goms.User/UpdateUser, 参数 {"Uid":"123","xxx":"name","Sex":"1"} 
grpcurl -plaintext -d '{"Uid":"123","xxx":"name","Sex":"1"}' localhost:50051 service.goms.User/UpdateUser

# 使用 grpc 方法 service.goms.User/ReadUser, 参数 {"Uid":"123"}
grpcurl -plaintext -d '{"Uid":"123"}' localhost:50051 service.goms.User/ReadUser

# 使用 grpc 方法 service.goms.User/DeleteUser, 参数 {"Uid":"123"}
grpcurl -plaintext -d '{"Uid":"123"}' localhost:50051 service.goms.User/DeleteUser
```