package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fuwensun/goms/eTest/internal/server/grpc"
	"github.com/fuwensun/goms/eTest/internal/server/http"
	"github.com/fuwensun/goms/eTest/internal/service"
)

type App struct {
	svc  *service.Service
	http *http.Server
	grpc *grpc.Server
}

func NewApp(svc *service.Service, h *http.Server, g *grpc.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
		http: h,
		grpc: g,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		log.Printf("server exit")
		fmt.Printf("context: %v\n", ctx)
		svc.Close()
		cancel()
	}
	return
}
