package http

import (
	"log"
	"net/http"

	"github.com/fuwensun/goms/eTest/internal/model"

	"github.com/gin-gonic/gin"
)

//
var pingcount model.PingCount

// example for http request handler.
func (srv *Server) ping(c *gin.Context) {
	svc := srv.svc

	pingcount++
	pc, err := svc.ReadHttpPingCount(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "read http ping count err",
		})
		return
	}
	err = svc.UpdateHttpPingCount(c, pingcount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update http ping count err",
		})
		return
	}
	msg := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   pc,
	})
	log.Printf("http ping msg: %v, count: %v", msg, pc)
}
