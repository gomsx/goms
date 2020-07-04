package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New.
func New() (engine *gin.Engine) {
	engine = gin.Default()
	initRouter(engine)
	engine.Run()
	return
}

// initRouter.
func initRouter(e *gin.Engine) {
	e.GET("/ping", ping)
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

