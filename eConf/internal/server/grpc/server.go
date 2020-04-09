package grpc

import (
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eConf/api"
	"github.com/fuwensun/goms/eConf/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var svc *service.Service

// Server.
type Server struct {
	// cfg *config
	gs  *grpc.Server
	svc *service.Service
}

// New.
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
