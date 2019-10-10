package grpc

import (
	// "fmt"
	// "context"
	"context"
	"log"
	"net"

	"github.com/fuwensun/goms/eGrpc/api"
	"github.com/fuwensun/goms/eGrpc/internal/service"
	xrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":7070"
)

var (
	svc *service.Service
)

//
type Server struct{}

//
func (s *Server) Ping(ctx context.Context, q *api.Request) (r *api.Reply, e error) {
	r = &api.Reply{Message: "pong" + " " + q.Message}
	log.Printf(r.Message)
	return r, nil
}

//
func New(s *service.Service) (server *Server) {
	log.Println("grpc new")
	svc = s

	server = &Server{} //server = new(Server)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	xs := xrpc.NewServer()
	api.RegisterApiServer(xs, server)
	reflection.Register(xs) //

	go func() {
		if err := xs.Serve(lis); err != nil {
			panic(err)
			// log.Fatalf("failed to serve: %v", err)
		}
	}()
	return
}
