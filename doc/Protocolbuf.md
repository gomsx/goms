# eApi

## Protocolbuf

### Protocol Compiler   

Protocol 编译器 protoc

```
# 安装
sudo apt install protobuf-compiler

# 查看
protoc --version

# 使用
protoc --help
```
>https://github.com/protocolbuffers/protobuf  
https://github.com/protocolbuffers/protobuf/releases  
https://developers.google.com/protocol-buffers/docs/overview 

### Protobuf Runtime
Go support for Protocol Buffers

```
# 安装
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# 查看
ls $GOPATH/src/github.com/golang
protobuf

ls $GOPATH/bin
protoc-gen-go

# 使用
protoc --go_out=:. *.proto
protoc --go_out=plugins=grpc:. *.proto
```
下载 protobuf 的 go 运行时，包含 protoc 编译器的插件和与之配对的包，如 protoc-gen-go 插件和 github.com/golang/protobuf/proto 包，还有 protoc-gen-go/grpc 插件和 google.golang.org/grpc 包(不在这个运行时中).并安装 protoc-gen-go 插件到 GOPATH 目录下, protoc-gen-go 提供　--go_out 参数，用于生成　go 代码，它包含静态子插件 grpc，提供　plugins=grpc　参数，用于生成　grpc 的桩代码.   

>https://github.com/golang/protobuf  
https://developers.google.com/protocol-buffers/docs/gotutorial  

### grpc Package 

这个包是 grpc 的 go 实现，已经映射为 github.com/grpc/grpc-go.

```
# 安装
go get -u google.golang.org/grpc

# 查看
ls $GOPATH/src/google.golang.org/grpc
ls $GOPATH/src/github.com/grpc/grpc-go
```

>https://github.com/grpc/grpc-go  
https://www.grpc.io/docs/tutorials/basic/go/   

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

