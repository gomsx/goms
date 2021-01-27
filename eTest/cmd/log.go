package main

import (
	m "github.com/aivuca/goms/eTest/internal/model"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infof("log level: %v", m.GetLogLevel())
}
