package main

import (
	"flag"
)

var (
	confpath string
)

func init() {
	flag.StringVar(&confpath, "conf", "../configs", "config path dir")
}

func flagx() {
	flag.Parse()
}
