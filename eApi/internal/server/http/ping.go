package http

import (
	"net/http"

	m "github.com/aivuca/goms/eApi/internal/model"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ping ping server.
func (s *Server) ping(ctx *gin.Context) {
	svc := s.svc
	//
	ping := &m.Ping{Type: "http"}
	ping, err := svc.HandPing(ctx, ping)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to hand ping, error: %v", err)
		return
	}
	//
	msg := ms.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   ping.Count,
	})
	log.Debugf("pong msg: %v, count: %v", msg, ping.Count)
	return
}
