package http

import (
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/fuwensun/goms/pkg/conf"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
		log.Warn().Msgf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Info().Msgf("get config file succ, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	//default
	cfg.Addr = ":8080"
	log.Info().Msgf("use default config, addr: %v", cfg.Addr)
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
	engine := gin.Default() // <==
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
	eng := s.eng
	go func() {
		if err := eng.Run(addr); err != nil {
			log.Fatal().Msgf("failed to run: %v", err)
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
	e.Use(setRequestId())
	//ping
	e.GET("/ping", s.ping)
	//user
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

// setRequestId set request id to request context.
func setRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set request_id
		setCtxRequestId(c)
		// before request
		c.Next()
	}
}

// setCtxRequestId gin.context With requestid.
func setCtxRequestId(ctx *gin.Context) {
	log.Debug().
		Msg("run request id middleware")
	id := ms.GetRequestId()
	c := ms.CarryCtxRequestId(ctx, id)
	ctx.Set("ctx", c)
	log.Ctx(c).Debug().
		Msg("new request id for new request")
}
