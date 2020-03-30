package dao

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("======> tear_up <======")
	tearupA()
	ret := m.Run()
	fmt.Println("======> tear_down <=======")
	teardownA()
	os.Exit(ret)
}

func tearupA() {
	command := "./testdata/tearA/up_docker.sh"
	// command := "ls -al"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}

func teardownA() {
	command := "./testdata/tearA/down_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}
