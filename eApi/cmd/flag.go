package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var (
	cfgpath string
)

func init() {
	flag.StringVar(&cfgpath, "cfgpath", "../configs", "config path")
}

//parseFlag parse cmd flag.
func parseFlag() {
	flag.Parse()
	log.Infof("config path: %v", cfgpath)
}
