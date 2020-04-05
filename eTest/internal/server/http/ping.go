package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ping
func (srv *Server) ping(c *gin.Context) {
	svc := srv.svc
	pc, err := svc.ReadHttpPingCount(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
		})
		return
	}
	pc++
	err = svc.UpdateHttpPingCount(c, pc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
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
