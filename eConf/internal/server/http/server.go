package http

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/aivuca/goms/pkg/conf"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/gin-gonic/gin"
)

// config config of server.
type config struct {
	Addr string `yaml:"addr"`
}

// Server server struct.
type Server struct {
	cfg *config
	eng *gin.Engine
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "http.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Printf("get config file error: %v", err)
	} else if cfg.Addr != "" {
		log.Printf("get config file succ, addr: %v", cfg.Addr)
		return cfg, nil
	}
	//TODO get env
	cfg.Addr = ":8080"
	log.Printf("use default config, addr: %v", cfg.Addr)
	return cfg, nil
}

// New new server and return.
func New(cfgpath string) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config: %v", err)
	}
	engine := gin.Default()
	server := &Server{
		cfg: cfg,
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

// initRouter init router.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
}

// ping ping server.
func ping(ctx *gin.Context) {
	msg := ms.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	log.Printf("pong msg: %v", msg)
	return
}
