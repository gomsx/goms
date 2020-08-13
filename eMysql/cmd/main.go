package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aivuca/goms/eMysql/internal/dao"
	"github.com/aivuca/goms/eMysql/internal/server/grpc"
	"github.com/aivuca/goms/eMysql/internal/server/http"
	"github.com/aivuca/goms/eMysql/internal/service"
)

func main() {
	fmt.Println("\n---eMysql---")
	parseFlag()

	dao := dao.New(cfgpath)
	log.Printf("new dao: %p", dao)

	svc := service.New(cfgpath, dao)
	log.Printf("new service: %p", svc)

	httpSrv := http.New(cfgpath, svc)
	log.Printf("new http server: %p", httpSrv)

	grpcSrv := grpc.New(cfgpath, svc)
	log.Printf("new grpc server: %p", grpcSrv)

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
