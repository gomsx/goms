package http

import (
	"net/http"

	m "github.com/fuwensun/goms/eApi/internal/model"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ping ping server.
func (s *Server) ping(ctx *gin.Context) {
	svc := s.svc
	c := ms.GetCtxVal(ctx)
	//
	ping := &m.Ping{Type: "http"}
	ping, err := svc.HandPing(c, ping)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().Msgf("failed to hand ping, error: %v", err)
		return
	}
	//
	msg := ms.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   ping.Count,
	})
	log.Ctx(c).Debug().Msgf("pong msg: %v, count: %v", msg, ping.Count)
	return
}
