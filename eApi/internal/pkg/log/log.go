package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lg = log.With().Str("service", "user.goms").Logger()

//
var Lg = lg

//
var Lgh = lg.With().Str("layer", "http server").Logger()

//
var Lgg = lg.With().Str("layer", "grpc server").Logger()

//
var Lgs = lg.With().Str("layer", "service").Logger()

//
var Lgd = lg.With().Str("layer", "dao").Logger()

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
