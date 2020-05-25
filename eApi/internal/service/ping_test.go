package service

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eApi/internal/dao/mock"
	. "github.com/fuwensun/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

//http
func TestHandPingHttp(t *testing.T) {
	Convey("TestHandPingHttp", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		Convey("for succ", func() {
			var pc PingCount = 2
			var want PingCount = 3
			daom.EXPECT().
				ReadPingCount(gomock.Any(), HTTP).
				Return(pc, nil)

			daom.EXPECT().
				UpdatePingCount(gomock.Any(), HTTP, pc+1).
				Return(nil)

			got, err := svc.HandPingHttp(context.Background())
			So(got, ShouldEqual, want)
			So(err, ShouldBeNil)
		})
	})
}

//grpc
func TestHandPingGrpc(t *testing.T) {
	Convey("TestHandPingGrpc", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		Convey("for succ", func() {
			var pc PingCount = 2
			var want PingCount = 3
			daom.EXPECT().
				ReadPingCount(gomock.Any(), GRPC).
				Return(pc, nil)

			daom.EXPECT().
				UpdatePingCount(gomock.Any(), GRPC, pc+1).
				Return(nil)

			got, err := svc.HandPingGrpc(context.Background())
			So(got, ShouldEqual, want)
			So(err, ShouldBeNil)
		})
	})
}
