package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server.
type Server struct {
	eng *gin.Engine
}

// New.
func New() *Server {
	engine := gin.Default()
	server := &Server{
		eng: engine,
	}
	initRouter(engine)

	addr := ":8080"
	go func() {
		if err := engine.Run(addr); err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()
	return server
}

// initRouter.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
	return
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
