package main

import (
	"flag"
	"fmt"
)

var (
	flagStr  string
	flagInt  int
	flagBool bool
)

func init() {

	flag.StringVar(&flagStr, "s", "", "")
	flag.IntVar(&flagInt, "i", 0, "a int value")
	flag.BoolVar(&flagBool, "bool", false, "")

	// flag.Parse()	//go run .  -s=string -i=1 -bool=true 会打印 usage
}

//
func parseFlag() {
	flag.Parse()

	fmt.Printf("\nflag command-line arguments:\ns = %v \ni = %v \nbool = %v\n\n",
		flagStr, flagInt, flagBool)

	fmt.Printf("non-flag command-line arguments all: %v\n", flag.Args())
	fmt.Printf("non-flag command-line arguments [0]: %v\n", flag.Arg(0))
}
