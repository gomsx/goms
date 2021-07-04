package http

import (
	"github.com/gomsx/goms/eApi/internal/service"
	"github.com/spf13/viper"

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
func getConfig() (*config, error) {
	cfg := &config{}
	//file
	cfg.Addr = viper.GetString("server.http.addr")
	if cfg.Addr != "" {
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
func New(s service.Svc) (*Server, error) {
	cfg, err := getConfig()
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
	log := v1.Group("/log")
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
