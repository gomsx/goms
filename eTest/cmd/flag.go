package main

import (
	"flag"

	"github.com/fuwensun/goms/eTest/internal/model"

	log "github.com/sirupsen/logrus"
)

var (
	cfgpath string
)

func init() {
	flag.StringVar(&cfgpath, "cfgpath", "../configs", "config path")
}

func parseFlag() {
	flag.Parse()
	log.Infof("config path: %v", cfgpath)
	model.CfgPath = cfgpath
}
