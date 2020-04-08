package grpc

import (
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eGrpc/api"
	"github.com/fuwensun/goms/eGrpc/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var svc *service.Service

//
type Server struct {
	// cfg *config
	gs  *grpc.Server
	svc *service.Service
}

//
func New(s *service.Service) *Server {
	gs := grpc.NewServer()
	server := &Server{
		// cfg: &cfg,
		svc: s,
		gs:  gs,
	}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	svc = s
	return server
}

// ping
func (s *Server) Ping(ctx context.Context, req *api.Request) (res *api.Reply, e error) {
	message := "pong" + " " + req.Message
	res = &api.Reply{Message: message}
	log.Printf("grpc" + " " + message)
	return res, nil
}
