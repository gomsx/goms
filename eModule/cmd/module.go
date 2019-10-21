package main

import (
	"github.com/fuwensun/goms/pkg/hello"
	"github.com/fuwensun/goms/pkg/world"
)

func module() {
	var h hello.Hello
	h.Print()
	var w world.World
	w.Print()
}
