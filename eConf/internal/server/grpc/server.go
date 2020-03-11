package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/api"
	"github.com/fuwensun/goms/eConf/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	xrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	svc      *service.Service
	conffile = "grpc.yml"
	addr     = ":50051"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

//
type Server struct{}

//
func New(s *service.Service) (server *Server) {
	svc = s

	var sc ServerConfig
	pathname := filepath.Join(svc.Confpath, conffile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Printf("get grpc server config file: %v", err)
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("grpc server addr: %v", addr)

	server = &Server{}

	lis, err := net.Listen("tcp", addr)
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
