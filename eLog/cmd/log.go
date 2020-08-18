package main

import (
	m "github.com/fuwensun/goms/eLog/internal/model"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msgf("log level = %v", m.GetLogLevel())
}
