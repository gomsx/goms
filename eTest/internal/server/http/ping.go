package http

import (
	"net/http"

	m "github.com/aivuca/goms/eTest/internal/model"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ping ping server.
func (s *Server) ping(ctx *gin.Context) {
	svc := s.svc
	c := getCtxVal(ctx)
	//
	ping := &m.Ping{Type: "http"}
	ping, err := svc.HandPing(ctx, ping)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Ctx(c).Info().
			Msgf("failed to hand ping, error: %v", err)
		return
	}
	//
	msg := ms.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   ping.Count,
	})
	log.Ctx(c).Debug().
		Msgf("pong msg: %v, count: %v", msg, ping.Count)
	return
}
