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

// example for grpc request handler.
func (s *Server) Ping(c context.Context, req *api.Request) (res *api.Reply, e error) {
	message := "pong" + " " + req.Message
	res = &api.Reply{Message: message}
	log.Printf("grpc" + " " + message)
	handping(c)
	return res, nil
}

//
var pc model.PingCount

//
func handping(c context.Context) {
	pc++
	svc.UpdateGrpcPingCount(c, pc)
	pc := svc.ReadGrpcPingCount(c)
	log.Printf("grpc ping count: %v\n", pc)
}
