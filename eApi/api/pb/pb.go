//go:generate protoc -I ./ -I ../../../pkg/proto --go_out=plugins=grpc:../ api.proto
//go:generate protoc -I ./ -I ../../../pkg/proto --grpc-gateway_out=logtostderr=true:../ api.proto
//go:generate protoc -I ./ -I ../../../pkg/proto --swagger_out=logtostderr=true:../ api.proto
package pb
