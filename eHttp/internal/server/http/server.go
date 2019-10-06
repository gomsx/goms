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
	g := e.Group("/test")
	{
		g.GET("/print", testPrint)
		g.GET("/ping", testPing)
	}
}

// example for http request handler.
func testPrint(c *gin.Context) {
	fmt.Println("http server ok!!!")
}
func testPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
