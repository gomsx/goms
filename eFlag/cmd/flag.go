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
	// flag.Usage = func() {
	// 	fmt.Printf("usage: xxx\n")
	// 	os.Exit(2)
	// }

	flag.StringVar(&flagStr, "s", "", "")
	flag.IntVar(&flagInt, "i", 0, "a int value")
	flag.BoolVar(&flagBool, "b", false, "")

	flag.Parse()
}

//go run .  -s=string -i=1 -b=true
//./cmd  -s=string -i=1 -b=true
func flagx() {
	fmt.Println("flagx()")
	args := flag.Args()
	fmt.Printf("args: %v", args)

	// var flagset flag.FlagSet
	// flagset.StringVar(&flagStr, "s", "", "")
	// flagset.IntVar(&flagInt, "i", 0, "int value")
	// flagset.BoolVar(&flagBool, "b", false, "")
	// flagset.Parse(args[1:])

	fmt.Printf("\nflag:\n s = %v\n i = %v\n b = %v\n\n", flagStr, flagInt, flagBool)
}
