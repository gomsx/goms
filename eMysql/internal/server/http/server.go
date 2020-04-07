package http

import (
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/model"
	"github.com/fuwensun/goms/eMysql/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

var svc *service.Service

type config struct {
	Addr string `yaml:"addr"`
}
type Server struct {
	cfg *config
	eng *gin.Engine
	svc *service.Service
}

func getConfig(cfgpath string) (config, error) {
	var cfg config
	filep := filepath.Join(cfgpath, "http.yml")
	if err := conf.GetConf(filep, &cfg); err != nil {
		log.Printf("get config file: %v", err)
	}
	if cfg.Addr != "" {
		log.Printf("get config addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":8080"
	log.Printf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

//
func New(cfgpath string, s *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panic(err)
	}
	engine := gin.Default()
	server := &Server{cfg: &cfg, eng: engine, svc: s}
	initRouter(engine)
	go func() {
		if err := engine.Run(cfg.Addr); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	svc = s
	return server
}

//
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
}

// ping
func ping(c *gin.Context) {
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)
	handping(c)
}

//
var pc model.PingCount

//
func handping(c *gin.Context) {
	pc++
	svc.UpdateHttpPingCount(c, pc)
	pc := svc.ReadHttpPingCount(c)
	log.Printf("http ping count: %v\n", pc)
}
