package http

import (
	"log"
	"net/http"

	m "github.com/aivuca/goms/eGrpc/internal/model"

	"github.com/gin-gonic/gin"
)

// Server server struct.
type Server struct {
	eng *gin.Engine
}

// New new server.
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

// initRouter init router.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
}

// ping ping methon.
func ping(c *gin.Context) {
	msg := m.MakePongMsg(c.Query("message"))
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	log.Printf("pong msg: %v", msg)
	return
}
