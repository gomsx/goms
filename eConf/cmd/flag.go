package main

import (
	"flag"
)

var (
	confpath string
	flagStr  string
	flagInt  int
	flagBool bool
)

func init() {
	flag.StringVar(&confpath, "conf", "../configs", "config path dir")
	// flag.IntVar(&flagInt, "i", 0, "a int value")
	// flag.BoolVar(&flagBool, "b", false, "")
}

//go run .  -s=string -i=1 -b=true
//./cmd  -s=string -i=1 -b=true
func flagx() {
	flag.Parse()

	// fmt.Println("flagx()")
	// args := flag.Args()
	// fmt.Printf("args: %v", args)

	// fmt.Printf("\nflag:\n s = %v\n i = %v\n b = %v\n\n", flagStr, flagInt, flagBool)
}
