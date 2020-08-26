package http

import (
	"log"
	"net/http"

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
	msg := makePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	log.Printf("pong msg: %v", msg)
	return
}

//makePongMsg make pong msg.
func makePongMsg(s string) string {
	if s == "" {
		s = "NONE!"
	}
	return "pong" + " " + s
}
