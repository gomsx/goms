package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/internal/pkg/conf"
	"github.com/fuwensun/goms/eConf/internal/service"

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
		log.Printf("get http server config file err: %v", err) //panic(err)
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("http server addr: %v", addr)

	engine = gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(addr); err != nil {
			log.Panicf("failed to serve: %v", err) //panic(err)
		}
	}()
	return
}

//
func initRouter(e *gin.Engine) {
	// e.GET("/ping", ping)
	callg := e.Group("/call")
	{
		callg.GET("/ping", ping)
	}
}

func ping(c *gin.Context) {
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)
}
