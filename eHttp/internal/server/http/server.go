package http

import (
	"fmt"

	"github.com/fuwensun/goms/eHttp/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *gin.Engine) {
	engine = gin.Default()
	initRouter(engine)
	engine.Run()
	return
}

//
func initRouter(e *gin.Engine) {
	// e.GET("/ping", ping)
	g := e.Group("/test")
	{
		g.GET("/ping", ping)
	}
}

// example for http request handler.
func ping(c *gin.Context) {
	fmt.Println("pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
