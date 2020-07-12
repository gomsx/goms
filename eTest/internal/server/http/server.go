package http

import (
	"path/filepath"

	"github.com/aivuca/goms/eTest/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
	//todo get env

	//default
	cfg.Addr = ":8080"
	log.Info().Msgf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

// New
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msg("get config, error")
		return nil, err
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default() // <==
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
	eng := srv.eng
	go func() {
		if err := eng.Run(addr); err != nil {
			log.Fatal().Msgf("failed to run: %v", err)
		}
	}()
}

func (srv *Server) Stop() {
	// h := srv.eng
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := h.Shutdown(ctx); err != nil {
	// 	log.Fatal().Msgf("Server forced to shutdown: %v", err)
	// }
	// log.Info().Msg("Server exiting")
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
