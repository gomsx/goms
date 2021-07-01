package grpc

import (
	"testing"

	api "github.com/gomsx/goms/eApi/api/v1"
	m "github.com/gomsx/goms/eApi/internal/model"
	"github.com/gomsx/goms/eApi/internal/service/mock"
	e "github.com/gomsx/goms/pkg/err"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	ping := &m.Ping{Type: "grpc"}
	want := &m.Ping{Type: "grpc", Count: 5}

	Convey("Ping with message", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, nil)
		//构建 req
		req := &api.PingReq{
			Data: &api.PingMsg{
				Message: "xxx",
			},
		}
		//发起 req
		res, err := srv.Ping(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Msg, ShouldEqual, "ok")
		So(res.Data.Count, ShouldEqual, want.Count)
		So(res.Data.Message, ShouldEqual, m.MakePongMsg(req.Data.Message))
	})

	Convey("Ping without message", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, nil)
		//构建 req
		req := &api.PingReq{}
		//发起 req
		res, err := srv.Ping(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Msg, ShouldEqual, "ok")
		So(res.Data.Count, ShouldEqual, want.Count)
		So(res.Data.Message, ShouldEqual, m.MakePongMsg(""))
	})

	Convey("Ping when service error", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, errx)
		//构建 req
		req := &api.PingReq{}
		//发起 req
		_, err := srv.Ping(ctxb, req)
		//断言
		So(err, ShouldEqual, errx)
	})
}
