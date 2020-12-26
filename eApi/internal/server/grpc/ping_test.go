package grpc

import (
	"errors"
	"testing"

	api "github.com/fuwensun/goms/eApi/api/v1"
	m "github.com/fuwensun/goms/eApi/internal/model"
	e "github.com/fuwensun/goms/eApi/internal/pkg/err"
	"github.com/fuwensun/goms/eApi/internal/service/mock"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	//
	srv := Server{svc: svcm}
	ctx := carryCtxRequestId(context.Background())
	errt := errors.New("error")
	ping := &m.Ping{Type: "grpc"}
	want := &m.Ping{Type: "grpc", Count: 5}

	Convey("TestPing should succ", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctx, ping).
			Return(want, nil)
		//构建 req
		req := &api.PingReq{
			Data: &api.PingMsg{
				Message: "xxx",
			},
		}
		//发起 req
		res, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Msg, ShouldEqual, "ok")
		So(res.Data.Count, ShouldEqual, want.Count)
		So(res.Data.Message, ShouldEqual, m.MakePongMsg(req.Data.Message))
	})

	Convey("TestPing should failed", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctx, ping).
			Return(want, errt)
		//构建 req
		req := &api.PingReq{}
		//发起 req
		_, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, errt)
	})
}
