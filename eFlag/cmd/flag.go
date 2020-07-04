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

	flag.StringVar(&flagStr, "s", "", "string value")
	flag.IntVar(&flagInt, "i", 0, "an int value")
	flag.BoolVar(&flagBool, "bool", false, "a bool value")
}

//
func parseFlag() {
	flag.Parse()

	fmt.Printf("\nflag command-line arguments:\ns = %v \ni = %v \nbool = %v\n\n",
		flagStr, flagInt, flagBool)

	fmt.Printf("non-flag command-line arguments all: %v\n", flag.Args())
	fmt.Printf("non-flag command-line arguments [0]: %v\n", flag.Arg(0))
	fmt.Printf("non-flag command-line arguments [0]: %v\n", flag.Arg(1))
}

