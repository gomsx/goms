package grpc

import (
	"log"
	"net"
	"path/filepath"

	"github.com/aivuca/goms/eRedis/api"
	"github.com/aivuca/goms/eRedis/internal/service"
	"github.com/aivuca/goms/pkg/conf"

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
	svc service.Svc
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	// file
	path := filepath.Join(cfgpath, "grpc.yaml")
	if err := conf.GetConf(path, cfg); err != nil {
		log.Printf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Printf("get config file succ, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":50051"
	log.Printf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		return nil, err
	}
	gs := grpc.NewServer()
	server := &Server{
		cfg: cfg,
		gs:  gs,
		svc: s,
	}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)
	return server, nil
}

// Start start server.
func (s *Server) Start() {
	addr := s.cfg.Addr
	gs := s.gs
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("tcp listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
}

// Stop stop server.
func (s *Server) Stop() {
	// todo
}
