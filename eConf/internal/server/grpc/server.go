package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/api"
	"github.com/fuwensun/goms/eConf/internal/pkg/conf"
	"github.com/fuwensun/goms/eConf/internal/service"

	xrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr    = ":7070"
	svc     *service.Service
	confile = "grpc.yml"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

//
type Server struct{}

//
func New(s *service.Service) (server *Server) {
	log.Println("grpc new")
	svc = s

	var sc ServerConfig
	pathname := filepath.Join(svc.Confpath, confile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		panic(err)
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}

	server = &Server{}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	xs := xrpc.NewServer()
	api.RegisterTestServer(xs, server)
	reflection.Register(xs) //

	go func() {
		if err := xs.Serve(lis); err != nil {
			panic(err)
			// log.Fatalf("failed to serve: %v", err)
		}
	}()
	return
}

// example for grpc request handler.
func (s *Server) Ping(ctx context.Context, q *api.Request) (r *api.Reply, e error) {
	r = &api.Reply{Message: "pong" + " " + q.Message}
	log.Printf(r.Message)
	return r, nil
}
