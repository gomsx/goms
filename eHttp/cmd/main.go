package main

import (
	"fmt"

	"github.com/fuwensun/goexample/eHttp/internal/server/http"
	"github.com/fuwensun/goexample/eHttp/internal/service"
)

func main() {
	fmt.Println("\n---eHttp---")
	fmt.Println("main()")

	// yamlx()
	flagx()

	svc := service.New()
	httpSrv := http.New(svc)
	fmt.Printf("%v", httpSrv)
}
