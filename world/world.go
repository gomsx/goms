package world

import "fmt"

type World struct{}

func (World) print() {
	fmt.Println("world\n")
}
