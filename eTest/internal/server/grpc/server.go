package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/api"
	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	xrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	svc        *service.Service
	cfgfile = "grpc.yml"
	addr       = ":50051"
)

type ServerCfg struct {
	Addr string `yaml:"addr"`
}

//
type Server struct{}

//
func New(cfgpath string, s *service.Service) (*Server, error) {
	svc = s

	var sc ServerCfg
	pathname := filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		err = fmt.Errorf("get config file: %w", err)
		return nil, err
	}
	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("grpc server addr: %v", addr)

	server := &Server{}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Errorf("tcp listen: %w", err)
		return server, err
	}
	xs := xrpc.NewServer()
	api.RegisterUserServer(xs, server)
	reflection.Register(xs)

	go func() {
		if err := xs.Serve(lis); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return server, nil
}

// example for grpc request handler.
func (s *Server) Ping(ctx context.Context, req *api.Request) (res *api.Reply, err error) {
	message := "pong" + " " + req.Message
	res = &api.Reply{Message: message}
	log.Printf("grpc" + " " + message)
	handping(ctx)
	return res, nil
}

//
var pingcount model.PingCount

//
func handping(c context.Context) {
	pingcount++
	svc.UpdateGrpcPingCount(c, pingcount)
	pc := svc.ReadGrpcPingCount(c)
	log.Printf("grpc ping count: %v\n", pc)
}
