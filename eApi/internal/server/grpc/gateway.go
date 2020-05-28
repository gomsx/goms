package grpc

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"

	"github.com/fuwensun/goms/eApi/api"
)

var mux = runtime.NewServeMux()

func newGateway(s *Server) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := api.RegisterUserHandlerServer(ctx, mux, s)
	if err != nil {
		return err
	}

	return nil
}

func xxxstart() {
	defer glog.Flush()
	if err := http.ListenAndServe(":8081", mux); err != nil {
		glog.Fatal(err)
	}
}
