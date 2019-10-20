package main

import (
	"flag"
	"log"
)

var (
	confpath string
)

func init() {
	flag.StringVar(&confpath, "conf", "../configs", "config path dir")
}

func parseFlag() {
	flag.Parse()
	log.Printf("config path: %v", confpath)
}
