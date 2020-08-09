package http

import (
	"path/filepath"
	"time"

	lg "github.com/fuwensun/goms/eApi/internal/pkg/log"
	rqid "github.com/fuwensun/goms/eApi/internal/pkg/requestid"
	"github.com/fuwensun/goms/eApi/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

// config.
type config struct {
	Addr string `yaml:"addr"`
}

// Server.
type Server struct {
	cfg *config
	eng *gin.Engine
	svc service.Svc
}

// log.
var log = lg.Lgh

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}

	//file
	filep := filepath.Join(cfgpath, "http.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Warn().Msg("get config file, error")
	}
	if cfg.Addr != "" {
		log.Info().Msgf("get config addr: %v", cfg.Addr)
		return cfg, nil
	}

	//env
	//get env
	//todo

	//default
	cfg.Addr = ":8080"
	log.Info().Msgf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New new http server and return.
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msg("get config, error")
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

// Start start http server.
func (srv *Server) Start() {
	addr := srv.cfg.Addr
	eng := srv.eng
	go func() {
		if err := eng.Run(addr); err != nil {
			log.Fatal().Msgf("failed to run: %v", err)
		}
	}()
}

// Stop stop http server.
func (srv *Server) Stop() {
	// h := srv.eng
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := h.Shutdown(ctx); err != nil {
	// 	log.Fatal().Msgf("Server forced to shutdown: %v", err)
	// }
	// log.Info().Msg("Server exiting")
}

// initRouter.
func (srv *Server) initRouter() {
	e := srv.eng
	//middleware
	e.Use(middlewarex())
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

// middlewarex.
func middlewarex() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()
		c.Next()
		// after request
		latency := time.Since(t)
		c.Set("latency", latency)
	}
}

// setRequestId set request id to request context.
func setRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set request_id
		id := rqid.Get()
		c.Set("request_id", id)
		log.Debug().Int64("request_id", id).Msg("new request id")
		// before request
		c.Next()
	}
}
