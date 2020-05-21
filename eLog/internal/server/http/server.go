package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eLog/internal/service"
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
func getConfig(cfgpath string) (config, error) {
	var cfg config
	filep := filepath.Join(cfgpath, "http.yml")
	if err := conf.GetConf(filep, &cfg); err != nil {
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
		cfg: &cfg,
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
	return
}

// initRouter
func (srv *Server) initRouter() {
	e := srv.eng
	e.GET("/ping", srv.ping)
	user := e.Group("/user")
	{
		user.POST("", srv.createUser)
		user.PUT("/:uid", srv.updateUser)
		user.GET("/:uid", srv.readUser)
		user.DELETE("/:uid", srv.deleteUser)
		user.GET("", srv.readUser)
	}
	return
}
