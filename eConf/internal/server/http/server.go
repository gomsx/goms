package http

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/fuwensun/goms/pkg/conf"
	"github.com/gin-gonic/gin"
)

// config
type config struct {
	Addr string `yaml:"addr"`
}

// Server.
type Server struct {
	cfg *config
	eng *gin.Engine
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
func New(cfgpath string) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to getConfig: %v", err)
	}
	engine := gin.Default()
	server := &Server{
		cfg: &cfg,
		eng: engine,
	}
	initRouter(engine)
	go func() {
		if err := engine.Run(cfg.Addr); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return server
}

// initRouter.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
}

// ping.
func ping(c *gin.Context) {
	msg := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	log.Printf("http ping msg: %v", msg)
	return
}
