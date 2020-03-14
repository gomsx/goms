package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

var (
	svc     service.Svc
	cfgfile = "http.yml"
	addr    = ":8080"
)

type ServerCfg struct {
	Addr string `yaml:"addr"`
}

type Server struct {
	eng *gin.Engine
	svc service.Svc
}

//
func New(cfgpath string, s service.Svc) (*Server, error) {
	var sc ServerCfg
	path := filepath.Join(cfgpath, cfgfile)
	if err := conf.GetConf(path, &sc); err != nil {
		log.Printf("get config file: %v", err)
		// fmt.Errorf("get config file: %w", err)
		// return nil, err
	}
	if sc.Addr != "" {
		addr = sc.Addr
	}
	log.Printf("http server addr: %v", addr)
	svc = s
	engine := gin.Default()
	server := &Server{eng: engine, svc: s}
	server.initRouter()
	server.start()
	return server, nil
}

//
func (srv *Server) initRouter() {
	e := srv.eng
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

//
func (srv *Server) start() {
	go func() {
		if err := srv.eng.Run(addr); err != nil {
			log.Panicf("failed to server: %v", err)
		}
	}()
}
