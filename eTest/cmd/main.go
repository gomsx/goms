package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fuwensun/goms/eTest/internal/app"
	"github.com/fuwensun/goms/eTest/internal/dao"
	"github.com/fuwensun/goms/eTest/internal/server/grpc"
	"github.com/fuwensun/goms/eTest/internal/server/http"
	"github.com/fuwensun/goms/eTest/internal/service"
)

func main() {
	fmt.Println("\n---eTest---")
	parseFlag()

	dao, cleandao, err := dao.New(cfgpath)
	if err != nil {
		panic(err)
	}

	svc, cleansvc, err := service.New(cfgpath, dao)
	if err != nil {
		cleandao()
		panic(err)
	}
	log.Printf("new service: %p", svc)

	httpSrv, err := http.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		panic(err)
	}
	log.Printf("http server start! addr: %p", httpSrv)

	grpcSrv, err := grpc.New(cfgpath, svc)
	if err != nil {
		cleansvc()
		cleandao()
		panic(err)
	}
	log.Printf("grpc server start! addr: %p", grpcSrv)

	app, cleanapp, err := app.NewApp(svc, httpSrv, grpcSrv)
	if err != nil {
		cleansvc()
		cleandao()
		panic(err)
	}
	log.Printf("new app: %p", app)

	app.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cleanapp()
			cleansvc()
			cleandao()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
