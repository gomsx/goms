package http

import (
	"context"
	"path/filepath"

	"github.com/fuwensun/goms/eApi/internal/service"
	"github.com/fuwensun/goms/pkg/conf"
	rqid "github.com/fuwensun/goms/pkg/requestid"

	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"
)

// config config of server.
type config struct {
	Addr string `yaml:"addr"`
}

// Server struct.
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
		log.Warn().Msgf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Info().Msgf("get config file, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//get env todo
	//default
	cfg.Addr = ":8080"
	log.Info().Msgf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msgf("get config error: %v", err)
		return nil, err
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default() //todo
	server := &Server{
		cfg: cfg,
		eng: engine,
		svc: s,
	}
	server.initRouter()

	log.Info().Msg("http server ok")
	return server, nil
}

// Start start server.
func (srv *Server) Start() {
	addr := srv.cfg.Addr
	eng := srv.eng
	go func() {
		if err := eng.Run(addr); err != nil {
			log.Fatal().Msgf("failed to run: %v", err)
		}
	}()
}

// Stop stop server.
func (srv *Server) Stop() {
}

// initRouter init router.
func (srv *Server) initRouter() {
	e := srv.eng
	//middleware
	e.Use(setRequestId())
	//group
	v1 := e.Group("/v1")
	//ping
	v1.GET("/ping", srv.ping)
	//log
	log := v1.Group("/logs")
	{
		log.GET("/:name", srv.readLog)   //Param
		log.PUT("/:name", srv.updateLog) //Param
		log.GET("", srv.readLog)         //Query
		log.PUT("", srv.updateLog)       //PostForm
	}
	//user
	users := v1.Group("/users")
	{
		users.POST("", srv.createUser)
		users.GET("/:uid", srv.readUser)
		users.PUT("/:uid", srv.updateUser)
		users.DELETE("/:uid", srv.deleteUser)
		users.GET("", srv.readUser)
		users.PUT("", srv.updateUser)
	}

}

// setRequestId set request id to request context.
func setRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set request_id
		log.Debug().Msg("run request id middleware")
		id := rqid.Get()
		lgx := log.With().Int64("request_id", id).Logger()
		ctx := lgx.WithContext(context.Background())
		c.Set("ctx", ctx)
		log.Debug().Int64("request_id", id).Msg("new request id for new request")
		// before request
		c.Next()
	}
}
