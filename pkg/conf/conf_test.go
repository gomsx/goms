package conf_test

import (
	"reflect"
	"testing"

	"github.com/fuwensun/goms/pkg/conf"
)

type TB struct {
	RenamedC int   `yaml:"c"`
	D        []int `yaml:",flow"`
}

type T struct {
	A string
	B TB
}

var want = T{
	A: "Easy!",
	B: TB{
		RenamedC: 2,
		D:        []int{3, 4},
	},
}

func TestConf(t *testing.T) {
	got := T{}
	if err := conf.GetConf("testData/yaml.yml", &got); err != nil {
		t.Fatalf("err: %v\n", err)
	}
	t.Logf("got = %v\n", got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}
