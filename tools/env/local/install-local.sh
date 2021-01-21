#!/bin/bash
set -x # set -e
set -u

# mysql redis
sudo apt install mysql-server -y
sudo apt install redis-server -y

# curl
sudo apt install curl

# protoc
sudo apt install protobuf-compiler

# protobuf 第一版
go get -u github.com/golang/protobuf
# protobuf 第二版 --> go get -u google.golang.org/protobuf
git clone https://github.com/protocolbuffers/protobuf-go.git $GOPATH/src/google.golang.org/protobuf

go get -u github.com/golang/protobuf/proto
go get -u google.golang.org/protobuf/cmd/protoc-gen-go

# grpc --> go get -u google.golang.org/grpc
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc

# grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2

## go-swagger
go get -u github.com/go-swagger/go-swagger/cmd/swagger

# grpcurl
go get -u github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl

# grpcui
go get -u github.com/fullstorydev/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcui

# mockgen
go get -u github.com/golang/mock/mockgen
go install github.com/golang/mock/mockgen
