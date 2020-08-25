package http

import (
	"log"
	"net/http"

	m "github.com/aivuca/goms/eRedis/internal/model"

	"github.com/gin-gonic/gin"
)

// ping ping server.
func (s *Server) ping(ctx *gin.Context) {
	svc := s.svc
	p := &m.Ping{Type: "http"}
	p, err := svc.HandPing(ctx, p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	//
	msg := m.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   p.Count,
	})
	log.Printf("pong msg: %v, count: %v", msg, p.Count)
	return
}
