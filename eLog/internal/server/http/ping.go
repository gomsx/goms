package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	. "github.com/fuwensun/goms/eLog/internal/model"
)

// ping
func (srv *Server) ping(c *gin.Context) {
	svc := srv.svc
	p := &Ping{}
	p.Type = "http"
	p, err := svc.HandPing(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	msg := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   p.Count,
	})
	log.Info().Msgf("http ping msg: %v, count: %v", msg, p.Count)
	return
}
