package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/aivuca/goms/eMysql/api"
	"github.com/aivuca/goms/eMysql/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var svc *service.Service

// config
type config struct {
	Addr string `yaml:"addr"`
}

// 问题：Server 依赖于 service.Service, 而它是个具体实现，违反了依赖倒置原则
// Server
type Server struct {
	cfg *config
	gs  *grpc.Server
	svc *service.Service
}

// getConfig
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	path := filepath.Join(cfgpath, "grpc.yml")
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

// New.
func New(cfgpath string, s *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config: %v", err)
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
	svc = s
	return server
}

// Ping
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	pc, err := svc.HandPingGrpc(c)
	if err != nil {
		res = &api.Reply{
			Message: "internal error!",
		}
		return res, err
	}
	msg := "pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		Count:   int64(pc),
	}
	log.Printf("grpc ping msg: %v, count: %v", msg, pc)
	return res, nil
}

