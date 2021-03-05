package main

import (
	"flag"

	"github.com/rs/zerolog/log"
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
	log.Info().Msgf("config path: %v", cfgpath)
}
