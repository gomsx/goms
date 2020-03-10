package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fuwensun/goms/eTest/internal/app"
	"github.com/fuwensun/goms/eTest/internal/server/grpc"
	"github.com/fuwensun/goms/eTest/internal/server/http"
	"github.com/fuwensun/goms/eTest/internal/service"
)

func main() {
	fmt.Println("\n---eTest---")
	parseFlag()

	svc := service.New(confpath)

	httpSrv := http.New(svc)
	log.Printf("http server start! addr: %v", &httpSrv)

	grpcSrv := grpc.New(svc)
	log.Printf("grpc server start! addr: %v", &grpcSrv)

	app, clean, err := app.NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		panic(err)
	}
	log.Printf("new app: %v", app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			clean()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
