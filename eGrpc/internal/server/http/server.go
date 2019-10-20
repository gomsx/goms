package http

import (
	"fmt"

	"github.com/fuwensun/goms/eGrpc/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service ////
)

//
func New(s *service.Service) (engine *gin.Engine) {
	svc = s
	engine = gin.Default()
	initRouter(engine)
	go func() {
		if err := engine.Run(); err != nil {
			panic(err) // log.Fatalf("failed to serve: %v", err)
		}
	}()
	return
}

//
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
	// g := e.Group("/test")
	// {
	// 	g.GET("/ping", ping)
	// }
}

// example for http request handler.
func ping(c *gin.Context) {
	fmt.Println("pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
