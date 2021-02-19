package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eRedis/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

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
	svc service.Svc
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "http.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Printf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Printf("get config file succeeded, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//TODO get env
	cfg.Addr = ":8080"
	log.Printf("use default config, addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
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

// Start start server.
func (s *Server) Start() {
	addr := s.cfg.Addr
	go func() {
		if err := s.eng.Run(addr); err != nil {
			log.Panicf("failed to server: %v", err)
		}
	}()
}

// Stop stop server.
func (s *Server) Stop() {
	// TODO
}

// initRouter init router.
func (s *Server) initRouter() {
	e := s.eng
	e.GET("/ping", s.ping)
	users := e.Group("/users")
	{
		users.POST("", s.createUser)
		users.GET("/:uid", s.readUser)
		users.PUT("/:uid", s.updateUser)
		users.DELETE("/:uid", s.deleteUser)
		users.GET("", s.readUser)
		users.PUT("", s.updateUser)
	}
}
