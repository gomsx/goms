package http

import (
	"net/http"

	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ping ping server.
func (s *Server) ping(ctx *gin.Context) {
	svc := s.svc
	c := getCtxVal(ctx)
	//
	p := &m.Ping{Type: "http"}
	p, err := svc.HandPing(c, p)
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
	log.Ctx(c).Debug().
		Msgf("pong msg: %v, count: %v", msg, p.Count)
	return
}
