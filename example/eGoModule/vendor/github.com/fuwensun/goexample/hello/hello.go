package hello

import "fmt"

type Hello struct{}

func (Hello) Print() {
	fmt.Println("hello\n")
}
