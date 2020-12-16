package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fuwensun/goms/eConf/internal/server/grpc"
	"github.com/fuwensun/goms/eConf/internal/server/http"
)

func main() {
	fmt.Println("\n---eConf---")
	parseFlag()

	httpSrv := http.New(cfgpath)
	log.Printf("new http server: %+v", httpSrv)

	grpcSrv := grpc.New(cfgpath)
	log.Printf("new grpc server: %+v", grpcSrv)

	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sch
		log.Printf("get a signal %v", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("server exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
