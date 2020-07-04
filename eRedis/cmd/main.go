package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aivuca/goms/eRedis/internal/app"
)

func main() {
	fmt.Println("\n---eRedis---")
	parseFlag()

	app, clean, err := app.InitApp(cfgpath)
	if err != nil {
		clean()
		panic(err)
	}
	app.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			clean()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

