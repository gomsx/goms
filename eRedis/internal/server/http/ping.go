package http

import (
	"log"

	"github.com/fuwensun/goms/eRedis/internal/model"

	"github.com/gin-gonic/gin"
)

// example for http request handler.
func ping(c *gin.Context) {
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)

	handping(c)
}

//
var pingcount model.PingCount

//
func handping(c *gin.Context) {
	pingcount++
	svc.UpdateHttpPingCount(c, pingcount)
	pc := svc.ReadHttpPingCount(c)
	log.Printf("http ping count: %v\n", pc)
}
