package dao

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var errx = errors.New("error xxx")
var ctxb = context.Background()

func getCfgPath() string {
	path := []string{
		"testdata/configs",
	}
	return path[0]
}
func TestMain(m *testing.M) {
	tearup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func tearup() {
	fmt.Println("==> tear_up")
	tearupEnv()
	tearupDb()
	tearupCache()
	tearupDao()
}

func teardown() {
	fmt.Println("==> tear_down")
	teardownDao()
	teardownCache()
	teardownDb()
	teardownEnv()
}

func tearupEnv() {
	command := "cd testdata/script && up_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: [ %s ] failed with error: %s\n", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell: [ %s ] succeed to finish with output:\n%s\n", command, string(output))
}

func teardownEnv() {
	command := "cd testdata/script && down_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: [ %s ] failed with error: %s\n", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell: [ %s ] succeed to finish with output:\n%s\n", command, string(output))
}
