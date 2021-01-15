package main

import (
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msgf("log level: %v", ms.GetLogLevel())
}
