package grpc

import (
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eGrpc/api"
	m "github.com/fuwensun/goms/eGrpc/internal/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server server struc.
type Server struct {
	gs *grpc.Server
}

// New new sever.
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

// Ping ping methon.
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	msg := m.MakePongMsg(req.Message)
	res = &api.Reply{
		Message: msg,
	}
	log.Printf("pong msg: %v", msg)
	return res, nil
}
