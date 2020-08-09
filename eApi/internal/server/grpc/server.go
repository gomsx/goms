package grpc

import (
	"net"
	"path/filepath"

	api "github.com/fuwensun/goms/eApi/api/v1"
	lg "github.com/fuwensun/goms/eApi/internal/pkg/log"
	rqid "github.com/fuwensun/goms/eApi/internal/pkg/requestid"
	"github.com/fuwensun/goms/eApi/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// config.
type config struct {
	Addr string `yaml:"addr"`
}

// Server.
type Server struct {
	cfg *config
	gs  *grpc.Server
	svc service.Svc
}

// log.
var log = lg.Lgg

// getConfig get config from file and env.
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

// New.
func New(cfgpath string, s service.Svc) (*Server, error) {
	//
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msg("get config, error")
		return nil, err
	}
	//
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(setRequestId()))
	gs := grpc.NewServer(opts...)
	//
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

// Start.
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

// Stop.
func (srv *Server) Stop() {
	//todo
}

func setRequestId() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = rqid.NewContext(ctx, rqid.Get())
		return handler(ctx, req)
	}
}
