package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eRedis/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

// config
type config struct {
	Addr string `yaml:"addr"`
}

// Server
type Server struct {
	cfg *config
	eng *gin.Engine
	svc service.Svc
}

// getConfig
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "http.yml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Printf("get config file: %v", err)
	}
	if cfg.Addr != "" {
		log.Printf("get config addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":8080"
	log.Printf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		return nil, err
	}
	engine := gin.Default()
	server := &Server{
		cfg: cfg,
		eng: engine,
		svc: s,
	}
	server.initRouter()
	return server, nil
}

// Start
func (srv *Server) Start() {
	addr := srv.cfg.Addr
	go func() {
		if err := srv.eng.Run(addr); err != nil {
			log.Panicf("failed to server: %v", err)
		}
	}()
}
func (srv *Server) Stop() {
	// ???
}

// initRouter
func (srv *Server) initRouter() {
	e := srv.eng
	e.GET("/ping", srv.ping)
	users := e.Group("/users")
	{
		users.POST("", srv.createUser)
		users.GET("/:uid", srv.readUser)
		users.PUT("/:uid", srv.updateUser)
		users.DELETE("/:uid", srv.deleteUser)
		users.GET("", srv.readUser)
		users.PUT("", srv.updateUser)
	}
}

