package main

import (
	"flag"
	"log"
)

var (
	confpath string
)

func init() {
	flag.StringVar(&confpath, "confpath", "../configs", "config path")
}

func parseFlag() {
	flag.Parse()
	log.Printf("config path: %v", confpath)
}
