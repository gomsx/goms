package http

import (
	"log"
	"net/http"

	m "github.com/fuwensun/goms/eMysql/internal/model"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
)

// ping ping server.
func (s *Server) ping(ctx *gin.Context) {
	svc := s.svc
	//
	ping := &m.Ping{Type: "http"}
	ping, err := svc.HandPing(ctx, ping)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	//
	msg := ms.MakePongMsg(ctx.Query("message"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   ping.Count,
	})
	log.Printf("pong msg: %v, count: %v", msg, ping.Count)
	return
}
