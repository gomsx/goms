package app

import (
	"github.com/gomsx/goms/eTest/internal/dao"
	"github.com/gomsx/goms/eTest/internal/server/grpc"
	"github.com/gomsx/goms/eTest/internal/server/http"
	"github.com/gomsx/goms/eTest/internal/service"

	log "github.com/sirupsen/logrus"
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
	log.Infof("==> 1, new dao: %p", dao)

	svc, cleansvc, err := service.New(cfgpath, dao)
	if err != nil {
		cleandao()
		return nil, nil, err
	}
	log.Infof("==> 2, new service: %p", svc)

	httpSrv, err := http.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Infof("==> 3, new http server: %p", httpSrv)

	grpcSrv, err := grpc.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Infof("==> 4, new grpc server: %p", grpcSrv)

	app, cleanapp, err := NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Infof("==> 5, new app: %p", app)

	return app, func() {
		cleanapp()
		cleansvc()
		cleandao()
	}, nil
}
