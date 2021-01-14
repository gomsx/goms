package main

import (
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msgf("log level: %v", ms.GetLogLevel())
}
