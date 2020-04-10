package http

import (
	"log"
	"net/http"

	."github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/fuwensun/goms/eRedis/internal/service"
	"github.com/gin-gonic/gin"
)

// ping
func (srv *Server) ping(c *gin.Context) {
	pc, err := handping(c, srv.svc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
		})
		return
	}
	msg := "pong" + " " + c.DefaultQuery("message", "NONE!")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"count":   pc,
	})
	log.Printf("http ping msg: %v, count: %v", msg, pc)
	return
}

// hangping
func handping(c *gin.Context, svc service.Svc) ( PingCount, error) {
	pc, err := svc.ReadHttpPingCount(c)
	if err != nil {
		return pc, err
	}
	pc++
	err = svc.UpdateHttpPingCount(c, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}
