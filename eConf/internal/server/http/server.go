package http

import (
	"log"
	"net/http"

	"github.com/fuwensun/goms/eConf/internal/service"

	"github.com/gin-gonic/gin"
)

var svc *service.Service
// Server.
type Server struct {
	// cfg *config
	eng *gin.Engine
	svc *service.Service
}

// New.
func New(s *service.Service) *Server {
	engine := gin.Default()
	server := &Server{
		// cfg: &cfg,
		eng: engine,
		svc: s,
	}
	initRouter(engine)
	go func() {
		if err := engine.Run(); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	svc = s
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
