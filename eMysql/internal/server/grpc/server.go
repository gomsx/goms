package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/aivuca/goms/eMysql/api"
	m "github.com/aivuca/goms/eMysql/internal/model"
	e "github.com/aivuca/goms/eMysql/internal/pkg/err"
	"github.com/aivuca/goms/eMysql/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// config config of server.
type config struct {
	Addr string `yaml:"addr"`
}

// Server server struc.
type Server struct {
	cfg *config
	gs  *grpc.Server
	svc *service.Service
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	path := filepath.Join(cfgpath, "grpc.yaml")
	if err := conf.GetConf(path, cfg); err != nil {
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

// New new sever.
func New(cfgpath string, s *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to getConfig: %v", err)
	}
	gs := grpc.NewServer()
	server := &Server{
		cfg: cfg,
		gs:  gs,
		svc: s,
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

// Ping ping methon.
func (s *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := s.svc
	//
	p := &m.Ping{}
	p.Type = "grpc"

	p, err := svc.HandPing(c, p)
	if err != nil {
		res = &api.Reply{
			Message: e.ErrInternalError.Error(),
		}
		return res, err
	}
	//
	res = &api.Reply{
		Message: m.MakePongMsg(req.Message),
		Count:   p.Count,
	}
	log.Printf("ping msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
