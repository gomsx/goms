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
	flag.BoolVar(&flagBool, "b", false, "")

	// flag.Parse()	//./cmd  -s=string -i=1 -b=true 会打印 usage
}

//go run .  -s=string -i=1 -b=true
//./cmd  -s=string -i=1 -b=true
func flagx() {
	flag.Parse()

	fmt.Println("flagx()")
	args := flag.Args()
	fmt.Printf("args: %v", args)

	fmt.Printf("\nflag:\n s = %v\n i = %v\n b = %v\n\n", flagStr, flagInt, flagBool)
}
