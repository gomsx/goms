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

>https://grpc-ecosystem.github.io/grpc-gateway/  
https://github.com/grpc-ecosystem/grpc-gateway　　

## swagger  

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
