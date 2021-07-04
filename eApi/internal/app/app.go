package app

import (
	"github.com/gomsx/goms/eApi/internal/dao"
	"github.com/gomsx/goms/eApi/internal/server/grpc"
	"github.com/gomsx/goms/eApi/internal/server/http"
	"github.com/gomsx/goms/eApi/internal/service"

	log "github.com/sirupsen/logrus"
)

// App.
type App struct {
	svc  service.Svc
	http *http.Server
	grpc *grpc.Server
}

// NewApp new app.
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
	log.Info("app ok")
	return
}

// Start start app.
func (app *App) Start() {
	app.http.Start()
	app.grpc.Start()
	return
}

// InitApp init app.
func InitApp() (*App, func(), error) {

	log.Infof("==> 1, new dao")
	dao, cleandao, err := dao.New()
	if err != nil {
		return nil, nil, err
	}
	log.Infof("dao ok")

	log.Infof("==> 2, new service")
	svc, cleansvc, err := service.New(dao)
	if err != nil {
		cleandao()
		return nil, nil, err
	}
	log.Infof("service ok")

	log.Infof("==> 3, new http server")
	httpSrv, err := http.New(svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Infof("http server ok")

	log.Infof("==> 4, new grpc server")
	grpcSrv, err := grpc.New(svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Infof("grpc server ok")

	log.Infof("==> 5, new app")
	app, cleanapp, err := NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Infof("app ok")

	return app, func() {
		cleanapp()
		cleansvc()
		cleandao()
	}, nil
}
