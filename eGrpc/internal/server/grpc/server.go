package grpc

import (
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eGrpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server.
type Server struct {
	gs *grpc.Server
}

// New.
func New() *Server {
	gs := grpc.NewServer()
	server := &Server{
		gs: gs,
	}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)

	port := ":50051"
	lis, err := net.Listen("tcp", port)
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

// Ping.
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	msg := "pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
	}
	log.Printf("grpc ping msg: %v", msg)
	return res, nil
}
