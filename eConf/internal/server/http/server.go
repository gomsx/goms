package http

import (
	"fmt"
	"path/filepath"

	"github.com/fuwensun/goms/eConf/internal/pkg/conf"
	"github.com/fuwensun/goms/eConf/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	svc     *service.Service
	confile = "http.yml"
	addr    = ":8080"
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

//
func New(s *service.Service) (engine *gin.Engine) {
	svc = s

	var sc ServerConfig
	pathname := filepath.Join(svc.Confpath, confile)
	if err := conf.GetConf(pathname, &sc); err != nil {
		panic(err)
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}

	engine = gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(addr); err != nil {
			panic(err) // log.Fatalf("failed to serve: %v", err)
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
