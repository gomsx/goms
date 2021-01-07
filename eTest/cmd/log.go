package main

import (
	m "github.com/aivuca/goms/eTest/internal/model"

	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msgf("log level: %v", m.GetLogLevel())
}
