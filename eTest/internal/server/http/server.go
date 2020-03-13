package http

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

var (
	svc        *service.Service
	cfgfile = "http.yml"
	addr       = ":8080"
)

type ServerCfg struct {
	Addr string `yaml:"addr"`
}

type Server = gin.Engine

//
func New(cfgpath string, s *service.Service) (*Server, error) {
	svc = s

	var sc ServerCfg
	path:= filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(path,&sc); err != nil {
		fmt.Errorf("get config file: %w", err)
		return nil, err
	}

	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("http server addr: %v", addr)

	engine := gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(addr); err != nil {
			log.Panicf("failed to server: %v", err)
		}
	}()
	return engine, nil
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
