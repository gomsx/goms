package http

import (
	"log"

	"github.com/fuwensun/goms/eGrpc/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service
)

//
func New(s *service.Service) (engine *gin.Engine) {
	svc = s
	engine = gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return
}

//
func initRouter(e *gin.Engine) {
	ug := e.Group("/user")
	{
		ug.GET("/ping", ping)
	}
}

func ping(c *gin.Context) {
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)
}
