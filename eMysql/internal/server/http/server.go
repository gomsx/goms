package http

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/service"
	"github.com/fuwensun/goms/pkg/conf"
	"github.com/gin-gonic/gin"
)

var svc *service.Service

// config
type config struct {
	Addr string `yaml:"addr"`
}

// Server.
type Server struct {
	cfg *config
	eng *gin.Engine
	svc *service.Service
}

// getConfig
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

// New.
func New(cfgpath string, s *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config: %v", err)
	}
	engine := gin.Default()
	server := &Server{
		cfg: &cfg,
		eng: engine,
		svc: s,
	}
	initRouter(engine)
	go func() {
		if err := engine.Run(cfg.Addr); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	svc = s
	return server
}

// initRouter.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
	return
}

// ping
func ping(c *gin.Context) {
	pc, err := svc.HandPingHttp(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
		})
		return
	}
	msg := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   pc,
	})
	log.Printf("http ping msg: %v, count: %v", msg, pc)
	return
}
