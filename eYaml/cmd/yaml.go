package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
	// "gopkg.in/yaml.v2"
	// "github.com/go-yaml/yaml"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3,4]
`

//
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func yamlx() {

	fmt.Println("yaml()")
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}
