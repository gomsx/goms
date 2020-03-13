package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/model"
	"github.com/fuwensun/goms/eMysql/internal/service"
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
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return
}

//
func initRouter(e *gin.Engine) {
	ug:= e.Group("/user")
	{
		ug.GET("/ping", ping)
	}
}

// example for http request handler.
func ping(c *gin.Context) {
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)
	handping(c)
}

//
var pingcount model.PingCount

//
func handping(c *gin.Context) {
	pingcount++
	svc.UpdateHttpPingCount(c, pingcount)
	pc := svc.ReadHttpPingCount(c)
	log.Printf("http ping count: %v\n", pc)
}
