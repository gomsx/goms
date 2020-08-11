package http

import (
	"log"
	"net/http"

	. "github.com/aivuca/goms/eRedis/internal/model"

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
	//
	msg := MakePongMsg(c.Query("message"))
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   p.Count,
	})
	log.Printf("ping msg: %v, count: %v", msg, p.Count)
	return
}
