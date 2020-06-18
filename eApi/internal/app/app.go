package app

import (
	"github.com/fuwensun/goms/eApi/internal/dao"
	. "github.com/fuwensun/goms/eApi/internal/pkg/log"
	"github.com/fuwensun/goms/eApi/internal/server/grpc"
	"github.com/fuwensun/goms/eApi/internal/server/http"
	"github.com/fuwensun/goms/eApi/internal/service"
)

//
type App struct {
	svc  service.Svc
	http *http.Server
	grpc *grpc.Server
}

//
var log = Lg

//
func NewApp(svc service.Svc, h *http.Server, g *grpc.Server) (app *App, close func(), err error) {
	app = &App{
		svc:  svc,
		http: h,
		grpc: g,
	}
	close = func() {
		h.Stop()
		g.Stop()
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
	log.Info().Msgf("==> 1, new dao: %p", dao)

	svc, cleansvc, err := service.New(cfgpath, dao)
	if err != nil {
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("==> 2, new service: %p", svc)

	httpSrv, err := http.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("==> 3, new http server: %p", httpSrv)

	grpcSrv, err := grpc.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("==> 4, new grpc server: %p", grpcSrv)

	app, cleanapp, err := NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("==> 5, new app: %p", app)

	return app, func() {
		cleanapp()
		cleansvc()
		cleandao()
	}, nil
}
