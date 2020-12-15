//go:generate protoc -I ./ -I ../../../../pkg/proto --go_out=plugins=grpc:../ api.proto
//go:generate protoc -I ./ -I ../../../../pkg/proto --grpc-gateway_out=logtostderr=true:../ api.proto
//go:generate protoc -I ./ -I ../../../../pkg/proto --openapiv2_out=logtostderr=true:../ api.proto
package pb

// 生成的代码中 "google/api" 要用 "google.golang.org/genproto/googleapis/api/annotations" 替代
// _ "google.golang.org/genproto/googleapis/api/annotations"
