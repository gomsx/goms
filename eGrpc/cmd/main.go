package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"os/signal"
	"syscall"

	"github.com/fuwensun/goexample/eGrpc/internal/server/http"
	"github.com/fuwensun/goexample/eGrpc/internal/service"
)

func main() {
	fmt.Println("\n---eGrpc---")
	fmt.Println("main()")

	// yamlx()
	// flagx()

	//
	svc := service.New()
	httpSrv := http.New(svc)
	fmt.Printf("%v\n", httpSrv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		// log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			// if err := httpSrv.Shutdown(ctx); err != nil {
			// log.Error("httpSrv.Shutdown error(%v)", err)
			// }
			// log.Info("server exit")
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
