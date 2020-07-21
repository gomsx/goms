package grpc

import (
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/api"
	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Addr string `yaml:"addr"`
}

//
type Server struct {
	cfg *config
	gs  *grpc.Server
	svc service.Svc
}

func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}

	//file
	path := filepath.Join(cfgpath, "grpc.yaml")
	if err := conf.GetConf(path, cfg); err != nil {
		log.Warn().Msg("get config file, error")
	}
	if cfg.Addr != "" {
		log.Info().Msgf("get config file, addr: %v", cfg.Addr)
		return cfg, nil
	}

	//env
	//todo get env

	//default
	cfg.Addr = ":50051"
	log.Info().Msgf("use default, addr: %v", cfg.Addr)
	return cfg, nil
}

//
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msg("get config, error")
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

func (s *Server) Start() {
	addr := s.cfg.Addr
	gs := s.gs
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Fatal().Msgf("failed to serve: %v", err)
		}
	}()
}

func (srv *Server) Stop() {
	// ???
}
