package api_test

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eMysql/api"
	"github.com/fuwensun/goms/eMysql/api/mock"
	"github.com/golang/mock/gomock"
)

func TestCallclient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCallClient(ctrl)
	in := &api.Request{Message: "ping"}
	want := &api.Reply{Message: "pong"}
	m.EXPECT().Ping(gomock.Any(), in).Return(want, nil)
	testCallClient(t, m, in, want.Message)
}

func testCallClient(t *testing.T, ec api.CallClient, in *api.Request, want string) {
	r, err := ec.Ping(context.Background(), in)
	if err != nil {
		t.Error("test failed:", err)
	}
	if got := r.Message; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
