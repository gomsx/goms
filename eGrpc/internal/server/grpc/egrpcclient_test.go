package grpc_test

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eGrpc/internal/server/grpc"
	mock "github.com/fuwensun/goms/eGrpc/internal/server/grpc/mock"
	"github.com/golang/mock/gomock"
)

func TestEgrpcclient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockEgrpcClient(ctrl)
	in := &grpc.Request{Message: "ping"}
	want := &grpc.Reply{Message: "pong"}
	m.EXPECT().Ping(gomock.Any(), in).Return(want, nil)
	testEgrpcclient(t, m, in, want.Message)
}

func testEgrpcclient(t *testing.T, ec grpc.EgrpcClient, in *grpc.Request, want string) {
	r, err := ec.Ping(context.Background(), in)
	if err != nil {
		t.Error("call failed:", err)
	}
	if got := r.Message; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
