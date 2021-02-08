package grpc

import (
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eGrpc/api"
	ms "github.com/fuwensun/goms/pkg/misc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server server struct.
type Server struct {
	gs *grpc.Server
}

// New new sever and return.
func New() *Server {
	gs := grpc.NewServer()
	server := &Server{
		gs: gs,
	}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)

	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return server
}

// Ping ping server.
func (s *Server) Ping(ctx context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	msg := ms.MakePongMsg(req.Message)
	res = &api.Reply{
		Message: msg,
	}
	log.Printf("pong msg: %v", msg)
	return res, nil
}
