package app

import (
	"github.com/aivuca/goms/eApi/internal/dao"
	"github.com/aivuca/goms/eApi/internal/server/grpc"
	"github.com/aivuca/goms/eApi/internal/server/http"
	"github.com/aivuca/goms/eApi/internal/service"

	"github.com/rs/zerolog/log"
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
	log.Info().Msg("app ok")
	return
}

// Start start app.
func (app *App) Start() {
	app.http.Start()
	app.grpc.Start()
	return
}

// InitApp init app.
func InitApp(cfgpath string) (*App, func(), error) {

	log.Info().Msgf("==> 1, new dao")
	dao, cleandao, err := dao.New(cfgpath)
	if err != nil {
		return nil, nil, err
	}
	log.Info().Msgf("dao ok")

	log.Info().Msgf("==> 2, new service")
	svc, cleansvc, err := service.New(cfgpath, dao)
	if err != nil {
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("service ok")

	log.Info().Msgf("==> 3, new http server")
	httpSrv, err := http.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("http server ok")

	log.Info().Msgf("==> 4, new grpc server")
	grpcSrv, err := grpc.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("grpc server ok")

	log.Info().Msgf("==> 5, new app")
	app, cleanapp, err := NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}
	log.Info().Msgf("app ok")

	return app, func() {
		cleanapp()
		cleansvc()
		cleandao()
	}, nil
}
