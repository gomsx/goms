package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.Output(os.Stderr)
	log.Logger = log.Level(zerolog.InfoLevel)
}
