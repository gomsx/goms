package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aivuca/goms/eApi/internal/app"

	"github.com/rs/zerolog/log"
)

func init() {
	fmt.Println("\n---eApi---")
}

func main() {
	parseFlag()

	log.Info().Msgf("app init ......")

	app, clean, err := app.InitApp(cfgpath)
	if err != nil {
		panic(err)
	}
	app.Start()

	log.Info().Msgf("app start ......")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info().Msgf("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			clean()
			log.Info().Msgf("app stop ......")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
