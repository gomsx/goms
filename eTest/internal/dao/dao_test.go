package dao

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/prashantv/gostub"
)

var ctx = context.Background()
var cfgpath = "testdata/configs"

func TestMain(m *testing.M) {
	cienv := os.Getenv("CI_ENV")
	if cienv != "travis" {
		fmt.Println("======> tear_up <======")
		tearupC()
	}
	ret := m.Run()
	if cienv != "travis" {
		fmt.Println("======> tear_down <=======")
		teardownC()
	}
	os.Exit(ret)
}

var cfgpathstub *gostub.Stubs

func tearupC() {
	//
	cfgpathstub = gostub.Stub(&cfgpath, "testdata/tearC/configs")
	fmt.Println(cfgpath)
	//
	command := "./testdata/tearC/up_docker.sh" // command := "ls -al"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}

func teardownC() {
	cfgpathstub.Reset()
	command := "./testdata/tearC/down_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}
