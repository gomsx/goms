package http

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

// ping
func (srv *Server) ping(c *gin.Context) {
	svc := srv.svc
	pc, err := svc.HandPingHttp(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	msg := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   pc,
	})
	log.Info().Msgf("http ping msg: %v, count: %v", msg, pc)
	return
}
