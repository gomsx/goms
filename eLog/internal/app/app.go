package app

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/fuwensun/goms/eLog/internal/dao"
	"github.com/fuwensun/goms/eLog/internal/server/grpc"
	"github.com/fuwensun/goms/eLog/internal/server/http"
	"github.com/fuwensun/goms/eLog/internal/service"
)

type App struct {
	svc  service.Svc
	http *http.Server
	grpc *grpc.Server
}

func NewApp(svc service.Svc, h *http.Server, g *grpc.Server) (app *App, close func(), err error) {
	app = &App{
		svc:  svc,
		http: h,
		grpc: g,
	}
	close = func() {
		//???
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		log.Info().Msgf("server exit")
		fmt.Printf("context: %v\n", ctx)
		cancel()
	}
	return
}

func (app *App) Start() {
	app.http.Start()
	app.grpc.Start()
	return
}

func InitApp(cfgpath string) (*App, func(), error) {
	dao, cleandao, err := dao.New(cfgpath)
	if err != nil {
		return nil, nil, err
	}
	log.Info().Msgf("new dao: %p", dao)

	svc, cleansvc, err := service.New(cfgpath, dao)
	if err != nil {
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("new service: %p", svc)

	httpSrv, err := http.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("new http server: %p", httpSrv)

	grpcSrv, err := grpc.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("new grpc server: %p", grpcSrv)

	app, cleanapp, err := NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("new app: %p", app)

	return app, func() {
		cleanapp()
		cleansvc()
		cleandao()
	}, nil
}
