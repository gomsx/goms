package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/fuwensun/goms/eTest/api"
	. "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service/mock"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestPing should succ", t, func() {

		//mock
		var pc PingCount = 2
		svcm.EXPECT().
			HandPingGrpc(gomock.Any()).
			Return(pc, nil)

		//构建 req
		var ctx = context.Background()
		req := &api.Request{
			Message: "xxx",
		}
		//发起 req
		resp, _ := srv.Ping(ctx, req)

		//断言
		So(resp.Message, ShouldEqual, "Pong xxx")
		So(resp.Count, ShouldEqual, pc)
	})

	Convey("TestPing should failed", t, func() {

		//mock
		err := errors.New("error")
		svcm.EXPECT().
			HandPingGrpc(gomock.Any()).
			Return(PingCount(0), err)

		//构建 req
		var ctx = context.Background()
		req := &api.Request{
			Message: "xxx",
		}
		//发起 req
		resp, _ := srv.Ping(ctx, req)

		//断言
		So(resp.Message, ShouldEqual, "internal error!")
	})
}
