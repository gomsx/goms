package http

import (
	"path/filepath"

	"github.com/aivuca/goms/eApi/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	//file
	filep := filepath.Join(cfgpath, "http.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Warnf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Infof("get config file succeeded, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//get env TODO
	//default
	cfg.Addr = ":8080"
	log.Infof("use default config, addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Errorf("get config error: %v", err)
		return nil, err
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default() //TODO
	server := &Server{
		cfg: cfg,
		eng: engine,
		svc: s,
	}
	server.initRouter()

	log.Info("http server ok")
	return server, nil
}

// Start start server.
func (s *Server) Start() {
	addr := s.cfg.Addr
	eng := s.eng
	go func() {
		if err := eng.Run(addr); err != nil {
			log.Fatalf("failed to run: %v", err)
		}
	}()
}

// Stop stop server.
func (s *Server) Stop() {
}

// initRouter init router.
func (s *Server) initRouter() {
	e := s.eng
	//middleware
	e.Use(middlewareX())
	//group
	v1 := e.Group("/v1")
	//ping
	v1.GET("/ping", s.ping)
	//log
	log := v1.Group("/logs")
	{
		log.GET("", s.readLog)   //Query
		log.PUT("", s.updateLog) //PostForm
	}
	//user
	users := v1.Group("/users")
	{
		users.POST("", s.createUser)
		users.GET("/:uid", s.readUser)   //Param
		users.PUT("/:uid", s.updateUser) //Param
		users.DELETE("/:uid", s.deleteUser)
		users.GET("", s.readUser)   //Query
		users.PUT("", s.updateUser) //PostForm
	}

}

// middlewareX.
func middlewareX() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}
