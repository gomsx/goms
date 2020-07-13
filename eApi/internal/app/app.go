package app

import (
	"path/filepath"

	"github.com/aivuca/goms/eApi/internal/dao"
	. "github.com/aivuca/goms/eApi/internal/pkg/log"
	"github.com/aivuca/goms/eApi/internal/server/grpc"
	"github.com/aivuca/goms/eApi/internal/server/http"
	"github.com/aivuca/goms/eApi/internal/service"
	"github.com/aivuca/goms/pkg/conf"
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
	log.Info().Msg("app ok")
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
		log.Warn().Msgf("get config file, %v", err)
	}
	if cfg.Ver != "" {
		log.Info().Msgf("get config file, ver: %v", cfg.Ver)
		return cfg, nil
	}
	//todo get env
	return cfg, nil
}

func InitApp(cfgpath string) (*App, func(), error) {

	_, err := getConfig(cfgpath)
	if err != nil {
		return nil, nil, err
	}

	log.Info().Msgf("==> 1, new dao")
	dao, cleandao, err := dao.New(cfgpath)
	if err != nil {
		return nil, nil, err
	}

	log.Info().Msgf("==> 2, new service")
	svc, cleansvc, err := service.New(cfgpath, dao)
	if err != nil {
		cleandao()
		return nil, nil, err
	}

	log.Info().Msgf("==> 3, new http server")
	httpSrv, err := http.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}

	log.Info().Msgf("==> 4, new grpc server")
	grpcSrv, err := grpc.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}

	log.Info().Msgf("==> 5, new app")
	app, cleanapp, err := NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		return nil, nil, err
	}

	return app, func() {
		cleanapp()
		cleansvc()
		cleandao()
	}, nil
}
