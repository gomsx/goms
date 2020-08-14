package dao

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

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
	fmt.Println("======> tear_up <======")
	tearup()
	ret := m.Run()
	fmt.Println("======> tear_down <=======")
	teardown()
	os.Exit(ret)
}

func tearup() {
	tearupDocker()
	tearupSqlmock()
}

func teardown() {
	teardownSqlmock()
	teardownDocker()
}

func tearupDocker() {
	if !isCiEnvDocker() {
		return
	}
	//
	command := "./testdata/teardocker/up_docker.sh" // command := "ls -al"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
	//等待 mysql docker 初始化完成
	time.Sleep(time.Second * 25)
}

func teardownDocker() {
	if !isCiEnvDocker() {
		return
	}
	command := "./testdata/teardocker/down_docker.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell: %s failed with error: %s\n", command, err.Error())
		// return
	}
	fmt.Printf("Execute Shell: %s finished with output:\n%s\n", command, string(output))
}
