package http

import (
	"log"
	"net/http"

	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
)

// Server server struct.
type Server struct {
	eng *gin.Engine
}

// New new server and return.
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

// ping ping server.
func ping(ctx *gin.Context) {
	msg := ms.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	log.Printf("pong msg: %v", msg)
	return
}
