package conf

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func GetConf(path string, data interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("---> file\n len = %v\n buf = %v\n\n", len(buf), buf)
	fmt.Printf("---> string\n%v\n\n", string(buf))

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("---> t:\n%v\n\n", data)
	return nil
}
