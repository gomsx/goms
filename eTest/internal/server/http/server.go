package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

var (
	svc        *service.Service
	configfile = "http.yml"
	addr       = ":8080"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type Server = gin.Engine

//
func New(s *service.Service) (engine *Server) {
	svc = s

	var sc ServerConfig
	pathname := filepath.Join(svc.Confpath, configfile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Printf("get http server config file: %v", err)
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("http server addr: %v", addr)

	engine = gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(addr); err != nil {
			log.Panicf("failed to server: %v", err)
		}
	}()
	return
}

//
func initRouter(e *gin.Engine) {

	e.GET("/ping", ping)
	user := e.Group("/user")
	{
		user.POST("", createUser)
		user.PUT("/:uid", updateUser)
		user.GET("/:uid", readUser)
		user.DELETE("/:uid", deleteUser)
		user.GET("/", readUser)
	}
}
