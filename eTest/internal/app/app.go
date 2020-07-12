package app

import (
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/aivuca/goms/eTest/internal/dao"
	"github.com/aivuca/goms/eTest/internal/server/grpc"
	"github.com/aivuca/goms/eTest/internal/server/http"
	"github.com/aivuca/goms/eTest/internal/service"
	"github.com/aivuca/goms/pkg/conf"
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

type config struct {
	Name string `yaml:"name"`
	Ver  string `yaml:"version"`
}

func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	//file
	path := filepath.Join(cfgpath, "app.yaml")
	if err := conf.GetConf(path, cfg); err != nil {
		log.Warn().Msgf("get config file, %w", err)
	}
	if cfg.Ver != "" {
		log.Info().Msgf("get config file, ver: %v", cfg.Ver)
		return cfg, nil
	}
	//env
	//todo get env
	return cfg, nil
}

func InitApp(cfgpath string) (*App, func(), error) {

	_, err := getConfig(cfgpath)
	if err != nil {
		return nil, nil, err
	}

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
