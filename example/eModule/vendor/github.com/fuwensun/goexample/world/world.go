package world

import "fmt"

type World struct{}

func (World) Print() {
	fmt.Println("world\n")
}
