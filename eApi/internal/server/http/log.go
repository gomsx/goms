package http

import (
	"net/http"

	lg "github.com/fuwensun/goms/eApi/internal/pkg/log"
	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

// readLog
func (srv *Server) readLog(c *gin.Context) {
	log.Debug().Msg("start to read log")

	nameStr := c.Param("name")
	if nameStr == "" {
		nameStr = c.Query("name")
	}
	name := "all" //todo
	log.Debug().Msgf("succ to create log date, name = %v", name)

	level := lg.GetLevel()

	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"level": level,
	})
	log.Debug().Msgf("succ to get log")
	return
}

// upateLog
func (srv *Server) updateLog(c *gin.Context) {
	log.Debug().Msg("start to update log")

	nameStr := c.Param("name")
	if nameStr == "" {
		nameStr = c.PostForm("name")
	}
	name := com.StrTo(nameStr).String()
	level := com.StrTo(c.PostForm("level")).String()
	log.Debug().Msgf("succ to create log date, name = %v, level = %v", name, level)

	lg.SetLevel(level)

	c.JSON(http.StatusOK, gin.H{})
	log.Debug().Msgf("succ to set log")
	return
}
