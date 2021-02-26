package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infof("log level: %v", log.GetLevel())
}
