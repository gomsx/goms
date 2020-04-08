package http

import (
	"log"

	"github.com/fuwensun/goms/eHttp/internal/service"

	"github.com/gin-gonic/gin"
)

var svc *service.Service

// New new a http server.
func New(s *service.Service) (engine *gin.Engine) {
	engine = gin.Default()
	initRouter(engine)
	engine.Run()
	return
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
