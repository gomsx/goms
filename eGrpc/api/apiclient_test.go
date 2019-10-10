package api_test

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eGrpc/api"
	"github.com/fuwensun/goms/eGrpc/api/mock"
	"github.com/golang/mock/gomock"
)

func TestApiclient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockApiClient(ctrl)
	in := &api.Request{Message: "ping"}
	want := &api.Reply{Message: "pong"}
	m.EXPECT().Ping(gomock.Any(), in).Return(want, nil)
	testApiClient(t, m, in, want.Message)
}

func testApiClient(t *testing.T, ec api.ApiClient, in *api.Request, want string) {
	r, err := ec.Ping(context.Background(), in)
	if err != nil {
		t.Error("call failed:", err)
	}
	if got := r.Message; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
