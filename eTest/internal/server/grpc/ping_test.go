package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/fuwensun/goms/eTest/api"
	m "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service/mock"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}
	//
	ctx := context.Background()
	errt := errors.New("error")
	ping := &m.Ping{Type: "grpc"}
	want := &m.Ping{Type: "grpc", Count: 5}

	Convey("TestPing should succ", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(gomock.Any(), ping).
			Return(want, nil)
		//构建 req
		req := &api.Request{
			Message: "xxx",
		}
		//发起 req
		res, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Message, ShouldEqual, m.MakePongMsg(req.Message))
		So(res.Count, ShouldEqual, want.Count)
	})

	Convey("TestPing should failed", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(gomock.Any(), ping).
			Return(want, errt)
		//构建 req
		req := &api.Request{}
		//发起 req
		_, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, errt)
	})
}
