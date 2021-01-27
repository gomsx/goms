package main

import (
	ms "github.com/aivuca/goms/pkg/misc"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infof("log level: %v", ms.GetLogLevel())
}
