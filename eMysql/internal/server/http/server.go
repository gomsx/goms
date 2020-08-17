package http

import (
	"log"
	"path/filepath"

	"github.com/aivuca/goms/eMysql/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

// config config of server.
type config struct {
	Addr string `yaml:"addr"`
}

// Server server struct.
type Server struct {
	cfg *config
	eng *gin.Engine
	svc *service.Service
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "http.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Printf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Printf("get config file, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":8080"
	log.Printf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string, svc *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to getConfig: %v", err)
	}
	engine := gin.Default()
	server := &Server{
		cfg: cfg,
		eng: engine,
		svc: svc,
	}
	initRouter(server, engine)
	go func() {
		if err := engine.Run(cfg.Addr); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return server
}

// initRouter init router.
func initRouter(s *Server, e *gin.Engine) {
	e.GET("/ping", s.ping)
}
