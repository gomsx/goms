package grpc

import (
	"net"
	"path/filepath"

	api "github.com/aivuca/goms/eApi/api/v1"
	lg "github.com/aivuca/goms/eApi/internal/pkg/log"
	"github.com/aivuca/goms/eApi/internal/service"
	"github.com/aivuca/goms/pkg/conf"

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

//
var log = lg.Lgg

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

	log.Info().Msg("grpc server ok")
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

	// gateway
	go func() {
		newGateway(s)
		// startGateway()
	}()
}

func (srv *Server) Stop() {
	// ???
}
