package http

import (
	"github.com/fuwensun/goexample/eHttp/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *gin.Engine) {
	engine = gin.Default()
	return
}
