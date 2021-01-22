## grpc-gateway

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

>https://grpc-ecosystem.github.io/grpc-gateway  
https://github.com/grpc-ecosystem/grpc-gateway  

