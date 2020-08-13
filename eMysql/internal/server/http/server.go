package http

import (
	"log"
	"net/http"
	"path/filepath"

	m "github.com/aivuca/goms/eMysql/internal/model"
	"github.com/aivuca/goms/eMysql/internal/service"
	"github.com/aivuca/goms/pkg/conf"

	"github.com/gin-gonic/gin"
)

// config config of server.
type config struct {
	Addr string `yaml:"addr"`
}

// Server server struc
type Server struct {
	cfg *config
	eng *gin.Engine
	svc *service.Service
}

// getConfig get config from file and env.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "http.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
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

// New new server.
func New(cfgpath string, svc *service.Service) *Server {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to getConfig: %v", err)
	}
	engine := gin.Default()
	server := &Server{
		cfg: cfg,
		eng: engine,
		svc: svc,
	}
	initRouter(server, engine)
	go func() {
		if err := engine.Run(cfg.Addr); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return server
}

// initRouter init router.
func initRouter(s *Server, e *gin.Engine) {
	e.GET("/ping", s.ping)
}

// ping ping methon.
func (s *Server) ping(c *gin.Context) {
	svc := s.svc
	//
	p := &m.Ping{}
	p.Type = "http"

	p, err := svc.HandPing(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	//
	msg := m.MakePongMsg(c.Query("message"))
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   p.Count,
	})
	log.Printf("ping msg: %v, count: %v", msg, p.Count)
	return
}
