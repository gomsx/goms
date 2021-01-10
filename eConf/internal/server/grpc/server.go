package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/api"
	"github.com/fuwensun/goms/pkg/conf"
	ms "github.com/fuwensun/goms/pkg/misc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// config config of server.
type config struct {
	Addr string `yaml:"addr"`
}

// Server server struct.
type Server struct {
	cfg *config
	gs  *grpc.Server
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	path := filepath.Join(cfgpath, "grpc.yaml")
	if err := conf.GetConf(path, cfg); err != nil {
		log.Printf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Printf("get config file succ, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":50051"
	log.Printf("use default config, addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config: %v", err)
	}
	gs := grpc.NewServer()
	server := &Server{
		cfg: cfg,
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

// Ping ping server.
func (s *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	msg := ms.MakePongMsg(req.Message)
	res = &api.Reply{
		Message: msg,
	}
	log.Printf("pong msg: %v", msg)
	return res, nil
}
