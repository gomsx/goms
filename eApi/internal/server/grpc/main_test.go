package grpc

import (
	"context"
	"fmt"
	"os"
	"testing"

	rqid "github.com/aivuca/goms/pkg/requestid"
	"github.com/rs/zerolog/log"
)

// var errx = errors.New("error")
var ctxx context.Context

func TestMain(m *testing.M) {
	fmt.Println("======> tear_up <======")
	tearup()
	ret := m.Run()
	fmt.Println("======> tear_down <=======")
	teardown()
	os.Exit(ret)
}
func tearup() {
	lgx := log.With().Int64("request_id", rqid.Get()).Logger()
	ctxx = lgx.WithContext(context.Background())
}

func teardown() {
}
