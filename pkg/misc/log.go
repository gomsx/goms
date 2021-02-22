package misc

import (
	log "github.com/sirupsen/logrus"
)

//
func GetLogLevel() string {
	level := log.GetLevel()
	return level.String()
}

//
func SetLogLevel(l string) {
	level, err := log.ParseLevel(l)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(log.Level(level))
}
