package http

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/pkg/conf"
	"github.com/fuwensun/goms/eMysql/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	svc     *service.Service
	conffile = "http.yml"
	addr    = ":8080"
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
	testg := e.Group("/test")
	{
		testg.GET("/ping", ping)
	}
}

// example for http request handler.
func ping(c *gin.Context) {
	fmt.Println("pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
