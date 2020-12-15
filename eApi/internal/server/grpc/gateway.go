package grpc

import (
	"net/http"

	api "github.com/fuwensun/goms/eApi/api/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// newGateway new gateway and start.
func newGateway(s *Server) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	// ok hander server
	// err := api.RegisterUserHandlerServer(ctx, mux, s)
	// ok endpoint
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterUserHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux)
}

// func startGateway() {
// 	defer glog.Flush()
// 	if err := http.ListenAndServe(":8081", mux); err != nil {
// 		glog.Fatal(err)
// 	}
// }
