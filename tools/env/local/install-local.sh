#!/bin/bash
set -x
set -e
set -u

# mysql redis
sudo apt install mysql-server
sudo apt install redis-server

# curl
sudo apt install curl

# grpcurl
go get -u github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl

# mockgen
go get -u github.com/golang/mock/mockgen
go install github.com/golang/mock/mockgen

# protoc
sudo apt install protobuf-compiler
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go install github.com/golang/protobuf/protoc-gen-go

# grpc
go get -u google.golang.org/grpc

# grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

