package model

import "github.com/rs/zerolog"

//
func GetLogLevel() string {
	level := zerolog.GlobalLevel()
	return level.String()
}

//
func SetLogLevel(l string) {
	level, err := zerolog.ParseLevel(l)
	if err != nil {
		level = zerolog.Level(zerolog.InfoLevel)
	}
	zerolog.SetGlobalLevel(zerolog.Level(level))
}
