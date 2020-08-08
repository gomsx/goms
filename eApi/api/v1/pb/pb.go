//go:generate protoc -I ./ -I ../../../../pkg/proto --go_out=plugins=grpc:../ api.proto
//go:generate protoc -I ./ -I ../../../../pkg/proto --grpc-gateway_out=logtostderr=true:../ api.proto
//go:generate protoc -I ./ -I ../../../../pkg/proto --swagger_out=logtostderr=true:../ api.proto
package pb

//生成的代码中 "google/api" 要用 "google.golang.org/genproto/googleapis/api/annotations" 替代

// import (
// 	context "context"
// 	fmt "fmt"

// 	// _ "google/api"
// 	math "math"

// 	proto "github.com/golang/protobuf/proto"
// 	empty "github.com/golang/protobuf/ptypes/empty"
// 	_ "google.golang.org/genproto/googleapis/api/annotations"
// 	grpc "google.golang.org/grpc"
// 	codes "google.golang.org/grpc/codes"
// 	status "google.golang.org/grpc/status"
// )
