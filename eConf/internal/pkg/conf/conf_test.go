package conf_test

import (
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eConf/internal/pkg/conf"

	"fmt"
)

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

type TT struct {
	RenamedC int   `yaml:"c"`
	D        []int `yaml:",flow"`
}

var want = T{
	A: "Easy!",
	B: TT{
		RenamedC: 2,
		D:        []int{3, 4},
	},
}

func TestConf(t *testing.T) {
	got := T{}
	if err := conf.GetConf("testData/yaml.yml", &got); err != nil {
		fmt.Printf("err: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}
