package main

import (
	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msgf("log level = %v", m.GetLogLevel())
}
