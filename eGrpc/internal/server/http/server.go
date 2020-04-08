package http

import (
	"log"

	"github.com/fuwensun/goms/eGrpc/internal/service"

	"github.com/gin-gonic/gin"
)

var svc *service.Service

type Server struct {
	// cfg *config
	eng *gin.Engine
	svc *service.Service
}

//
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

//
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
}

func ping(c *gin.Context) {
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)
}
