#!/bin/bash
set -ex
set -u

sudo apt install mysql
sudo apt install redis
sudo apt install curl

#grpcurl
go get -u github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
#mockgen
go get -u github.com/golang/mock/mockgen
go install github.com/golang/mock/mockgen
#protoc
go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
