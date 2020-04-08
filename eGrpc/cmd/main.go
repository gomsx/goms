package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fuwensun/goms/eGrpc/internal/server/grpc"
	"github.com/fuwensun/goms/eGrpc/internal/server/http"
	"github.com/fuwensun/goms/eGrpc/internal/service"
)

func main() {
	fmt.Println("\n---eGrpc---")

	svc := service.New()
	log.Printf("new service: %p", svc)

	httpSrv := http.New(svc)
	log.Printf("new http server: %p", httpSrv)

	grpcSrv := grpc.New(svc)
	log.Printf("new grpc server: %p", grpcSrv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
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
