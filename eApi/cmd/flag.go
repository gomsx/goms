package main

import (
	"flag"

	"github.com/aivuca/goms/eApi/internal/model"

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
	model.CfgPath = cfgpath
}
