package http

import (
	"net/http"

	. "github.com/fuwensun/goms/eApi/internal/model"
	"github.com/fuwensun/goms/eApi/internal/pkg/reqid"

	"github.com/gin-gonic/gin"
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
	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("ping msg: %v, count: %v", msg, p.Count)
	return
}
