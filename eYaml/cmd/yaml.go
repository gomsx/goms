package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2" // "github.com/go-yaml/yaml"
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

func handyaml() {

	//to struct
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	//to map
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	//hand yaml file
	buf, err := ioutil.ReadFile("./configs/yaml.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--> f dump:\n%v\n\n", string(buf))

	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--> t:\n%v\n\n", t)
}
