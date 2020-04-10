package main

import (
	"flag"
	"log"

	"github.com/fuwensun/goms/eRedis/internal/model"
)

var (
	cfgpath string
)

func init() {
	flag.StringVar(&cfgpath, "cfgpath", "../configs", "config path")
}

func parseFlag() {
	flag.Parse()
	log.Printf("config path: %v", cfgpath)
	model.CfgPath = cfgpath
}
