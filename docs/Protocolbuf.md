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
# 第一版
go get -u github.com/golang/protobuf

# 第二版/修定版
## 国际标准安装
go get -u google.golang.org/protobuf

## 国内从 github 镜像安装
git clone https://github.com/protocolbuffers/protobuf-go.git $GOPATH/src/google.golang.org/protobuf

```

>This project is comprised of two components:  
Code generator: The protoc-gen-go tool is a compiler plugin to protoc, the protocol buffer compiler. It augments the protoc compiler so that it knows how to generate Go specific code for a given .proto file.  
Runtime library: The protobuf module contains a set of Go packages that form the runtime implementation of protobufs in Go. This provides the set of interfaces that define what a message is and functionality to serialize message in various formats (e.g., wire, JSON, and text).

插件
```
# 安装
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/protobuf/cmd/protoc-gen-go 

# 查看
ls $GOPATH/bin
protoc-gen-go

# 使用
protoc --go_out=:. *.proto
protoc --go_out=plugins=grpc:. *.proto
```
下载 protobuf 的 go 运行时，包含 protoc 编译器的插件和与之配对的包，如 protoc-gen-go 插件和 github.com/golang/protobuf/proto 包，还有 protoc-gen-go/grpc 插件和 google.golang.org/grpc 包(不在这个运行时中).并安装 protoc-gen-go 插件到 GOPATH 目录下, protoc-gen-go 提供　--go_out 参数，用于生成　go 代码，它包含静态子插件 grpc，提供　plugins=grpc　参数，用于生成　grpc 的桩代码.   

>https://github.com/golang/protobuf  
https://pkg.go.dev/github.com/golang/protobuf?readme=expanded#section-readme  
https://github.com/protocolbuffers/protobuf-go  
https://pkg.go.dev/google.golang.org/protobuf?readme=expanded#section-readme  
https://developers.google.com/protocol-buffers/docs/gotutorial  

### grpc Package

这个包是 grpc 的 go 实现.

```
# 国际标准安装
go get -u google.golang.org/grpc

# 国内从 github 镜像安装
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc

# 查看
ls $GOPATH/src/google.golang.org/grpc
```

>https://github.com/grpc/grpc-go  
https://www.grpc.io/docs/tutorials/basic/go/  

### grpc-gateway

```
# 安装
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2

# 查看
ls $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
protoc-gen-grpc-gateway
protoc-gen-openapiv2

ls $GOPATH/bin
protoc-gen-grpc-gateway
protoc-gen-openapiv2

# 使用
protoc --grpc-gateway_out=logtostderr=true:. *.proto
protoc --swagger_out=logtostderr=true:. *.proto
```

>https://grpc-ecosystem.github.io/grpc-gateway/  
https://github.com/grpc-ecosystem/grpc-gateway　　

