package hello

import "fmt"

type Hello struct{}

func (Hello) Print() string {
	s := "hello\n"
	fmt.Println(s)
	return s
}
