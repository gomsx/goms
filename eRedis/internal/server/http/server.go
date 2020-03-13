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
	cfgfile = "http.yml"
	addr     = ":8080"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

//
func New(s *service.Service) (engine *gin.Engine) {
	svc = s

	var sc ServerConfig
	pathname := filepath.Join(svc.Cfgpath, cfgfile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		log.Printf("get http server config file! error: %v", err)
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

	e.GET("/ping", ping)
	user := e.Group("/user")
	{
		user.POST("", createUser)
		user.PUT("/:uid", updateUser)
		user.GET("/:uid", readUser)
		user.DELETE("/:uid", deleteUser)
		user.GET("", readUser)
	}
}
