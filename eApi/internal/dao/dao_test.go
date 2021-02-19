package dao

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var errx = fmt.Errorf("error xxx")
var ctxb = context.Background()

func isCiEnvDocker() bool {
	ciEnvDocker := os.Getenv("CI_ENV_DOCKER")
	fmt.Printf("CI_ENV_DOCKER == %v\n", ciEnvDocker)
	if ciEnvDocker == "no" || ciEnvDocker == "" {
		return false
	}
	return true
}

func getCfgPath() string {
	path := []string{
		"testdata/configs",
		"testdata/teardocker/configs",
	}
	if isCiEnvDocker() {
		return path[1]
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
	tearupDocker()
	tearupDb()
	tearupCache()
	tearupDao()
}

func teardown() {
	fmt.Println("==> tear_down")
	teardownDao()
	teardownCache()
	teardownDb()
	teardownDocker()
}

func tearupDocker() {
	if !isCiEnvDocker() {
		return
	}
	command := "./testdata/teardocker/up_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: [ %s ] failed with error: %s\n", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell: [ %s ] succeed to finish with output:\n%s\n", command, string(output))
}

func teardownDocker() {
	if !isCiEnvDocker() {
		return
	}
	command := "./testdata/teardocker/down_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: [ %s ] failed with error: %s\n", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell: [ %s ] succeed to finish with output:\n%s\n", command, string(output))
}
