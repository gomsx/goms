package misc

import (
	"context"

	"github.com/gin-gonic/gin"
)

// GetCtxVal get context val from gin.Context.
func GetCtxVal(ctx *gin.Context) context.Context {
	return ctx.MustGet("ctx").(context.Context)
}
