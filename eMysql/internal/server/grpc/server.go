package grpc

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/api"
	"github.com/fuwensun/goms/eMysql/internal/model"
	"github.com/fuwensun/goms/eMysql/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var svc *service.Service

type config struct {
	Addr string `yaml:"addr"`
}

//
type Server struct {
	cfg *config
	gs  *grpc.Server
	svc *service.Service
}

func getConfig(cfgpath string) (config, error) {
	var cfg config
	path := filepath.Join(cfgpath, "grpc.yml")
	if err := conf.GetConf(path, &cfg); err != nil {
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

func New(cfgpath string, s *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panic(err)
	}
	gs := grpc.NewServer()
	server := &Server{cfg: &cfg, svc: s, gs: gs}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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
	pc, err := handping(c, svc)
	if err != nil {
		res = &api.Reply{
			Message: "internal error!",
		}
		return res, err
	}
	msg := "pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		// Count:   pc,
	}
	log.Printf("grpc ping msg: %v count: %v", msg, pc)
	return res, nil
}

// hangping
func handping(c context.Context, svc *service.Service) (model.PingCount, error) {
	pc, err := svc.ReadGrpcPingCount(c)
	if err != nil {
		return pc, err
	}
	pc++
	err = svc.UpdateGrpcPingCount(c, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}
