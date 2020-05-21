package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fuwensun/goms/eLog/internal/app"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(os.Stdout)
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
func main() {
	fmt.Println("\n---eLog---")
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
		log.Info().Msgf("get a signal: %s", s.String())
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
