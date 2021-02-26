package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

// readLog
func (s *Server) readLog(ctx *gin.Context) {
	log.Debug("start to read log")

	level := log.GetLevel().String()
	ctx.JSON(http.StatusOK, gin.H{
		"level": level,
	})

	log.Debugf("succeeded to get log level: %v", level)
	return
}

// upateLog
func (s *Server) updateLog(ctx *gin.Context) {
	log.Debug("start to update log")

	lv := com.StrTo(ctx.PostForm("level")).String()
	log.Debugf("succeeded to create log data level: %v", lv)

	level, err := log.ParseLevel(lv)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	ctx.JSON(http.StatusOK, gin.H{})

	log.Debugf("succeeded to set log level: %v", level.String())
	return
}
