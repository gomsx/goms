package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/api"
	"github.com/fuwensun/goms/pkg/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// config
type config struct {
	Addr string `yaml:"addr"`
}

// Server.
type Server struct {
	cfg *config
	gs  *grpc.Server
}

// getConfig
func getConfig(cfgpath string) (config, error) {
	var cfg config
	path := filepath.Join(cfgpath, "grpc.yml")
	if err := conf.GetConf(path, &cfg); err != nil {
		log.Printf("get config file: %v", err)
	}
	if cfg.Addr != "" {
		log.Printf("get config addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":50051"
	log.Printf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New.
func New(cfgpath string) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to getConfig: %v", err)
	}
	gs := grpc.NewServer()
	server := &Server{
		cfg: &cfg,
		gs:  gs,
	}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)

	lis, err := net.Listen("tcp", cfg.Addr)
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
