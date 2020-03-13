package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fuwensun/goms/eRedis/internal/server/grpc"
	"github.com/fuwensun/goms/eRedis/internal/server/http"
	"github.com/fuwensun/goms/eRedis/internal/service"
)

func main() {
	fmt.Println("\n---eRedis---")
	parseFlag()

	svc := service.New(cfgpath)

	httpSrv := http.New(svc)
	log.Printf("http server start! addr: %p", httpSrv)

	grpcSrv := grpc.New(svc)
	log.Printf("grpc server start! addr: %p", grpcSrv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			log.Printf("server exit")
			fmt.Printf("context: %v\n", ctx)
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
