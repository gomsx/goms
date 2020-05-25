package grpc

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eApi/api"
	. "github.com/fuwensun/goms/eApi/internal/model"
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
		var pc PingCount = 2
		svcm.EXPECT().
			HandPingGrpc(gomock.Any()).
			Return(pc, nil)

		//构建 req
		req := &api.Request{
			Message: "xxx",
		}
		//发起 req
		resp, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(resp.Message, ShouldEqual, "Pong xxx")
		So(resp.Count, ShouldEqual, pc)
	})

	Convey("TestPing should failed", t, func() {
		//mock
		svcm.EXPECT().
			HandPingGrpc(gomock.Any()).
			Return(PingCount(0), ErrInternalError)

		//构建 req
		req := &api.Request{}
		//发起 req
		_, err := srv.Ping(ctx, req)
		//断言
		So(err, ShouldEqual, ErrInternalError)
	})
}
