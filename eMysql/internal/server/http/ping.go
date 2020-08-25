package http

import (
	"log"
	"net/http"

	m "github.com/aivuca/goms/eMysql/internal/model"

	"github.com/gin-gonic/gin"
)

// ping ping methon.
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
	log.Printf("pong msg: %v, count: %v", msg, p.Count)
	return
}
