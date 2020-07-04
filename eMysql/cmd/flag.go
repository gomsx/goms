package main

import (
	"flag"
	"log"
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
}

