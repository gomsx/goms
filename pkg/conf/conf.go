package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//
func GetConf(path string, data interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("failed to readfile! : %w", err)
		return err

	}
	err = yaml.Unmarshal(buf, data)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal! : %w", err)
		return err
	}
	return nil
}