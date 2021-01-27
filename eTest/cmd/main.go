package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aivuca/goms/eTest/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("\n---eTest---")
	parseFlag()

	log.Infof("app init ......")

	app, clean, err := app.InitApp(cfgpath)
	if err != nil {
		panic(err)
	}
	app.Start()

	log.Infof("app start ......")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			clean()
			log.Infof("app stop ......")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
