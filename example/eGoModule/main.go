package main // import "github.com/fuwensun/e1"
import (
	"fmt"

	"github.com/fuwensun/goexample/hello"
	"github.com/fuwensun/goexample/world"
)

func main() {
	var h hello.Hello
	h.Print()
	var w world.World
	w.Print()
	fmt.Println("hehe\n")
}
