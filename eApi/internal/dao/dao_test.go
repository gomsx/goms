package dao

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/prashantv/gostub"
)

var cfgpath = "testdata/configs"
var ctx = context.Background()

func isCiEnvDockers() bool {
	ciEnvDocker := os.Getenv("CI_ENV_DOCKER")
	fmt.Printf("CI_ENV_DOCKER == %v\n", ciEnvDocker)
	if ciEnvDocker == "no" || ciEnvDocker == "" {
		return false
	}
	return true
}
func TestMain(m *testing.M) {
	fmt.Println("======> tear_up <======")
	tearup()
	ret := m.Run()
	fmt.Println("======> tear_down <=======")
	teardown()
	os.Exit(ret)
}

var cfgstub *gostub.Stubs

func tearup() {
	cfgstub = gostub.Stub(&cfgpath, "testdata/configs")
	fmt.Printf("stub config path to: %v", cfgpath)
	tearupdocker()
	tearupSqlmock()
}

func teardown() {
	teardownSqlmock()
	teardowndocker()
	cfgstub.Reset()
}

func tearupdocker() {
	if !isCiEnvDockers() {
		return
	}
	cfgstub = gostub.Stub(&cfgpath, "testdata/teardocker/configs")
	fmt.Printf("stub config path to: %v", cfgpath)
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

func teardowndocker() {
	if !isCiEnvDockers() {
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
	cfgstub.Reset()
}
