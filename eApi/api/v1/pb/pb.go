
//
//go:generate protoc -I ./ -I ../../../../pkg/proto --go_out=:.. --go-grpc_out=require_unimplemented_servers=false:.. api.proto
//go:generate protoc -I ./ -I ../../../../pkg/proto --grpc-gateway_out=logtostderr=true:../ api.proto
//go:generate protoc -I ./ -I ../../../../pkg/proto --openapiv2_out=logtostderr=true:../ api.proto
package pb

// doc

// < 1 >
// 生成的代码中 "google/api" 要用下面的替代
// _ "google.golang.org/genproto/googleapis/api/annotations"

// < 2.1 >
// https://github.com/grpc/grpc-go
// protoc -I ./ -I ../../../../pkg/proto --go_out=plugins=grpc:../ api.proto

// https://github.com/grpc-ecosystem/grpc-gateway
// protoc -I ./ -I ../../../../pkg/proto --go_out ./gen/go/ --go_opt paths=source_relative --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative api.proto
// protoc -I ./ -I ../../../../pkg/proto --go_out=:.. --go-grpc_out=require_unimplemented_servers=false:.. api.proto

// https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc

// < 2.2>
// go-protobuf-V1 ==> github.com/golang/protobuf == https://github.com/golang/protobuf
// go-protobuf-V2 ==> google.golang.org/protobuf == https://github.com/protocolbuffers/protobuf-go
// grpc ==> google.golang.org/grpc == https://github.com/grpc/grpc-go.git

// https://github.com/golang/protobuf/commit/cea45d6ceb8761b4bed31f8826a76722c053b3ad#diff-af5ff51dd13c7b1ca8f83edd0b7c1d26f79a2a501a2e01cc2b783751a9e53e19
// --> protoc-gen-go: move gengogrpc into v1 repo
// --> The eventual home of this is the gRPC repo, but extract it from the APIv2 repo for now.
// --> gengogrpc: V1 --> V2 --> grpc

