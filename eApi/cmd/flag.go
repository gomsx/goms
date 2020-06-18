package main

import (
	"flag"

	"github.com/fuwensun/goms/eApi/internal/model"
)

var (
	cfgpath string
)

func init() {
	flag.StringVar(&cfgpath, "cfgpath", "../configs", "config path")
}

func parseFlag() {
	flag.Parse()
	log.Info().Msgf("config path: %v", cfgpath)
	model.CfgPath = cfgpath
}
