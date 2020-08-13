package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New server.
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

// ping ping methon.
func ping(c *gin.Context) {
	msg := makePongMsg(c.Query("message"))
	c.JSON(http.StatusOK, gin.H{
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
