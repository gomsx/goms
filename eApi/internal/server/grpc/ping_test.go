package grpc

import (
	"context"
	"testing"

	api "github.com/fuwensun/goms/eApi/api/v1"
	m "github.com/fuwensun/goms/eApi/internal/model"
	e "github.com/fuwensun/goms/eApi/internal/pkg/err"
	"github.com/fuwensun/goms/eApi/internal/service/mock"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}
	var ctx = context.Background()

	Convey("TestPing should succ", t, func() {
		//mock
		p := &m.Ping{
			Type: "grpc",
		}
		want := &m.Ping{
			Type:  "grpc",
			Count: 5,
		}
		svcm.EXPECT().
			HandPing(gomock.Any(), p).
			Return(want, nil)

		//构建 req
		req := &api.Request{
			Message: "xxx",
		}
		//发起 req
		resp, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(resp.Message, ShouldEqual, "Pong xxx")
		So(resp.Count, ShouldEqual, want.Count)
	})

	Convey("TestPing should failed", t, func() {
		//mock
		p := &m.Ping{
			Type: "grpc",
		}
		want := &m.Ping{
			Type:  "grpc",
			Count: 5,
		}
		svcm.EXPECT().
			HandPing(gomock.Any(), p).
			Return(want, e.ErrInternalError)

		//构建 req
		req := &api.Request{}
		//发起 req
		_, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}
