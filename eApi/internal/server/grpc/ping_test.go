package grpc

import (
	"testing"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	"github.com/aivuca/goms/eApi/internal/service/mock"
	"golang.org/x/net/context"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	ctx := ctxCarryRqid(context.Background())
	//
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestPing should succ", t, func() {
		ping := &m.Ping{
			Type: "grpc",
		}
		want := &m.Ping{
			Type:  "grpc",
			Count: 5,
		}
		//mock
		svcm.EXPECT().
			HandPing(ctx, ping).
			Return(want, nil)

		//构建 req
		data := &api.PingMsg{
			Message: "xxx",
		}
		req := &api.PingReq{
			Data: data,
		}
		//发起 req
		res, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Msg, ShouldEqual, "ok")
		So(res.Data.Count, ShouldEqual, want.Count)
		So(res.Data.Message, ShouldEqual, m.MakePongMsg(data.Message))
	})

	Convey("TestPing should failed", t, func() {
		ping := &m.Ping{
			Type: "grpc",
		}
		want := &m.Ping{
			Type:  "grpc",
			Count: 5,
		}
		//mock
		svcm.EXPECT().
			HandPing(ctx, ping).
			Return(want, e.ErrInternalError)

		//构建 req
		req := &api.PingReq{}
		//发起 req
		_, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}
