package http

import (
	"net/http"

	m "github.com/fuwensun/goms/eTest/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ping ping server.
func (s *Server) ping(c *gin.Context) {
	svc := s.svc
	//
	p := &m.Ping{Type: "http"}
	p, err := svc.HandPing(c, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	//
	msg := m.MakePongMsg(c.Query("message"))
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   p.Count,
	})
	log.Ctx(c).Debug().
		Msgf("pong msg: %v, count: %v", msg, p.Count)
	return
}
