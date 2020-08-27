package http

import (
	"net/http"

	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/unknwon/com"
)

// readLog
func (s *Server) readLog(ctx *gin.Context) {
	log.Debug().Msg("start to read log")

	name := com.StrTo(ctx.Param("name")).String()
	if name == "" {
		name = com.StrTo(ctx.Query("name")).String()
	}

	name = "all" //todo
	log.Debug().
		Msgf("succ to create log data, name = %v", name)

	level := m.GetLogLevel()

	ctx.JSON(http.StatusOK, gin.H{
		"name":  name,
		"level": level,
	})
	log.Debug().
		Msgf("succ to get log, name = %v, level = %v", name, level)
	return
}

// upateLog
func (s *Server) updateLog(ctx *gin.Context) {
	log.Debug().Msg("start to update log")

	name := com.StrTo(ctx.Param("name")).String()
	if name == "" {
		name = com.StrTo(ctx.PostForm("name")).String()
	}
	level := com.StrTo(ctx.PostForm("level")).String()

	log.Debug().
		Msgf("succ to create log data, name = %v, level = %v", name, level)

	m.SetLogLevel(level)

	ctx.JSON(http.StatusOK, gin.H{})
	log.Debug().
		Msgf("succ to set log, name = %v, level = %v", name, level)
	return
}
