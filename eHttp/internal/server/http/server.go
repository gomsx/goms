package http

import (
	"log"
	"net/http"

	m "github.com/gomsx/goms/eHttp/internal/model"

	"github.com/gin-gonic/gin"
)

// New server and return.
func New() (engine *gin.Engine) {
	engine = gin.Default()
	initRouter(engine)
	go engine.Run()
	return
}

// initRouter init router.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
}

// ping ping server.
func ping(ctx *gin.Context) {
	msg := m.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	log.Printf("pong msg: %v", msg)
	return
}
