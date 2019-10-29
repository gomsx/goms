package api_test

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eMysql/api"
	"github.com/fuwensun/goms/eMysql/api/mock"
	"github.com/golang/mock/gomock"
)

func TestTestclient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockTestClient(ctrl)
	in := &api.Request{Message: "ping"}
	want := &api.Reply{Message: "pong"}
	m.EXPECT().Ping(gomock.Any(), in).Return(want, nil)
	testTesClient(t, m, in, want.Message)
}

func testTesClient(t *testing.T, ec api.TestClient, in *api.Request, want string) {
	r, err := ec.Ping(context.Background(), in)
	if err != nil {
		t.Error("call failed:", err)
	}
	if got := r.Message; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
