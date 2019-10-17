package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"os/signal"
	"syscall"

	"github.com/fuwensun/goms/eConf/internal/server/grpc"
	"github.com/fuwensun/goms/eConf/internal/server/http"
	"github.com/fuwensun/goms/eConf/internal/service"
)

func main() {
	fmt.Println("\n---eConf---")
	fmt.Println("main()")

	// yamlx()
	// flagx()

	//
	svc := service.New()

	httpSrv := http.New(svc)
	log.Printf("http server start : %v\n", httpSrv)

	grpcSrv := grpc.New(svc)
	log.Printf("grpc server start : %v\n", grpcSrv)

	log.Printf("=== server start !!! ===")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			// if err := httpSrv.Shutdown(ctx); err != nil {
			// log.Error("httpSrv.Shutdown error(%v)", err)
			// }
			log.Printf("server exit")
			fmt.Printf("context: %v\n", ctx) //
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
