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

//
func GetLevel() string {
	level := zerolog.GlobalLevel()
	return level.String()
}

//
func SetLevel(l string) {
	level, err := zerolog.ParseLevel(l)
	if err != nil {
		level = zerolog.Level(zerolog.InfoLevel)
	}
	zerolog.SetGlobalLevel(zerolog.Level(level))
}

//
func GetLevelInt() int64 {
	level := zerolog.GlobalLevel()
	return int64(level)
}

//
func SetLevelInt(l int64) {
	zerolog.SetGlobalLevel(zerolog.Level(l))
}
