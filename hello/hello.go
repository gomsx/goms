package hello

import "fmt"

type Hello struct{}

func (Hello) print() {
	fmt.Println("hello\n")
}
