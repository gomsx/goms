package world

import "fmt"

type World struct{}

func (World) Print() string {
	s := "world\n"
	fmt.Println(s)
	return s
}
