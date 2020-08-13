package grpc

import (
	"net"
	"path/filepath"

	api "github.com/aivuca/goms/eApi/api/v1"
	lg "github.com/aivuca/goms/eApi/internal/pkg/log"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"
	"github.com/aivuca/goms/eApi/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"golang.org/x/net/context"
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

// New server.
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

	// gateway
	go func() {
		newGateway(s)
		// startGateway()
	}()
}

// Stop stop server.
func (s *Server) Stop() {
	//todo
}

// setRequestId set request id to context.
func setRequestId() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		id := rqid.Get()
		ctx = rqid.NewContext(ctx, id)
		log.Debug().Int64("request_id", id).Msg("new request id for new request")
		return handler(ctx, req)
	}
}
