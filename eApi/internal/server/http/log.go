package http

import (
	"net/http"

	m "github.com/fuwensun/goms/eApi/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/unknwon/com"
)

// readLog
func (srv *Server) readLog(ctx *gin.Context) {
	// c := ctx.MustGet("ctx").(context.Context)
	log.Debug().Msg("start to read log")

	nameStr := ctx.Param("name")
	if nameStr == "" {
		nameStr = ctx.Query("name")
	}
	name := "all" //todo
	log.Debug().Msgf("succ to create log date, name = %v", name)

	level := m.GetLogLevel()

	ctx.JSON(http.StatusOK, gin.H{
		"name":  name,
		"level": level,
	})
	log.Debug().Msgf("succ to get log")
	return
}

// upateLog
func (srv *Server) updateLog(ctx *gin.Context) {
	// c := ctx.MustGet("ctx").(context.Context)
	log.Debug().Msg("start to update log")

	nameStr := ctx.Param("name")
	if nameStr == "" {
		nameStr = ctx.PostForm("name")
	}
	name := com.StrTo(nameStr).String()
	level := com.StrTo(ctx.PostForm("level")).String()
	log.Debug().Msgf("succ to create log date, name = %v, level = %v", name, level)

	m.SetLogLevel(level)

	ctx.JSON(http.StatusOK, gin.H{})
	log.Debug().Msgf("succ to set log")
	return
}
