package grpc

import (
	"context"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eLog/api"
	"github.com/fuwensun/goms/eLog/internal/service"
	"github.com/fuwensun/goms/pkg/conf"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/rs/zerolog/log"
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
		log.Warn().Msgf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Info().Msgf("get config file succ, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	//default
	cfg.Addr = ":50051"
	log.Info().Msgf("use default config, addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string, svc service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msgf("get config error: %v", err)
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
		svc: svc,
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
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Fatal().Msgf("failed to serve: %v", err)
		}
	}()
}

// Stop stop server.
func (s *Server) Stop() {
	//todo
}

// setRequestId set request id to context.
func setRequestId() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		c := carryCtxRequestId(ctx)
		return handler(c, req)
	}
}

// carryCtxRequestId context carry requestid.
func carryCtxRequestId(ctx context.Context) context.Context {
	ctx = log.Logger.WithContext(ctx)
	id := ms.GetRequestId()
	return ms.CarryCtxId(ctx, "request_id", id)
}
