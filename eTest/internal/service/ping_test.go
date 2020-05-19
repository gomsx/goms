package service

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eTest/internal/dao/mock"
	. "github.com/fuwensun/goms/eTest/internal/model"

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

// func TestUpdateGrpcPingCount(t *testing.T) {
// 	Convey("TestUpdateGrpcPingCount should return nil", t, func() {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()
// 		daom := mock.NewMockDao(ctrl)
// 		svc := service{dao: daom}

// 		var pc PingCount = 2
// 		daom.EXPECT().
// 			UpdatePingCount(gomock.Any(), GRPC, pc).
// 			Return(nil)

// 		err := svc.UpdateGrpcPingCount(context.Background(), pc)
// 		So(err, ShouldBeNil)
// 	})
// }

// func TestReadGrpcPingCount(t *testing.T) {
// 	Convey("TestReadGrpcPingCount ", t, func() {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()
// 		daom := mock.NewMockDao(ctrl)
// 		svc := service{dao: daom}

// 		Convey("for succ", func() {
// 			var want PingCount = 2
// 			daom.EXPECT().
// 				ReadPingCount(gomock.Any(), GRPC).
// 				Return(want, nil)

// 			got, err := svc.ReadGrpcPingCount(context.Background())
// 			So(got, ShouldEqual, want)
// 			So(err, ShouldBeNil)
// 		})
// 	})
// }
