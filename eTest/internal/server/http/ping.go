package http

import (
	"log"

	"github.com/fuwensun/goms/eTest/internal/model"

	"github.com/gin-gonic/gin"
)

//
var pingcount model.PingCount

// example for http request handler.
func (srv *Server) ping(c *gin.Context) {
	svc := srv.svc
	message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(200, gin.H{
		"message": message,
	})
	log.Printf("http" + " " + message)

	pingcount++
	svc.UpdateHttpPingCount(c, pingcount)
	pc := svc.ReadHttpPingCount(c)
	log.Printf("http ping count: %v\n", pc)
}
