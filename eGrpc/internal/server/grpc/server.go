package grpc

import (
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eGrpc/api"
	"github.com/fuwensun/goms/eGrpc/internal/service"

	xrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	svc  *service.Service
	port = ":50051"
)

//
type Server struct{}

//
func New(s *service.Service) (server *Server) {
	svc = s

	server = &Server{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	xs := xrpc.NewServer()
	api.RegisterUserServer(xs, server)
	reflection.Register(xs)

	go func() {
		if err := xs.Serve(lis); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return
}

// example for grpc request handler.
func (s *Server) Ping(ctx context.Context, q *api.Request) (r *api.Reply, e error) {
	message := "pong" + " " + q.Message
	r = &api.Reply{Message: message}
	log.Printf("grpc" + " " + message)
	return r, nil
}
