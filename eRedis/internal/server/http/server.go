package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eRedis/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

var (
	svc      *service.Service
	conffile = "http.yml"
	addr     = ":8080"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

//
func New(s *service.Service) (engine *gin.Engine) {
	svc = s

	var sc ServerConfig
	pathname := filepath.Join(svc.Confpath, conffile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Printf("failed to get http server config file! error: %v", err)
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("http server addr: %v", addr)

	engine = gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(addr); err != nil {
			log.Panicf("failed to serve! error: %v", err)
		}
	}()
	return
}

//
func initRouter(e *gin.Engine) {
	callg := e.Group("/call")
	{
		callg.GET("/ping", ping)
	}

	userg := e.Group("/user")
	{
		userg.GET("/updatename", updatename)
	}
}
