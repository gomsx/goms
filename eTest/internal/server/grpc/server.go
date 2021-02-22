package grpc

import (
	"net"
	"path/filepath"

	"github.com/aivuca/goms/eTest/api"
	"github.com/aivuca/goms/eTest/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	log "github.com/sirupsen/logrus"
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
	//file
	path := filepath.Join(cfgpath, "grpc.yaml")
	if err := conf.GetConf(path, cfg); err != nil {
		log.Warnf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Infof("get config file succeeded, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//TODO get env
	//default
	cfg.Addr = ":50051"
	log.Infof("use default config, addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Errorf("get config error: %v", err)
		return nil, err
	}
	//
	gs := grpc.NewServer()
	//
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
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

// Stop stop server.
func (s *Server) Stop() {
	//TODO
}
