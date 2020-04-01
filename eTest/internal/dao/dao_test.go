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
	fmt.Println("======> tear_up <======")
	// tearupA()
	ret := m.Run()
	fmt.Println("======> tear_down <=======")
	// teardownA()
	os.Exit(ret)
}

var cfgpathstub *gostub.Stubs

func tearupA() {
	//
	cfgpathstub = gostub.Stub(&cfgpath, "testdata/tearA/configs")
	//
	command := "./testdata/tearA/up_docker.sh" // command := "ls -al"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}

func teardownA() {
	cfgpathstub.Reset()
	command := "./testdata/tearA/down_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}
