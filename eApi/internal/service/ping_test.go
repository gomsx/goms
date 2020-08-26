package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fuwensun/goms/eApi/internal/dao/mock"
	m "github.com/fuwensun/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandPing(t *testing.T) {
	Convey("TestHandPing", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		dao := mock.NewMockDao(ctrl)
		//
		svc := service{dao: dao}
		ctx := context.Background()
		errt := errors.New("error")

		Convey("for succ", func() {
			ping := &m.Ping{Type: "http", Count: 2}
			want := &m.Ping{Type: "http", Count: 3}

			dao.EXPECT().
				ReadPing(ctx, ping.Type).
				Return(ping, nil)

			dao.EXPECT().
				UpdatePing(ctx, ping).
				Return(nil)

			got, err := svc.HandPing(ctx, ping)

			So(err, ShouldBeNil)
			So(reflect.DeepEqual(got, want), ShouldBeTrue)
		})

		Convey("for failed", func() {
			ping := &m.Ping{Type: "http", Count: 2}

			dao.EXPECT().
				ReadPing(ctx, ping.Type).
				Return(ping, errt)

			got, err := svc.HandPing(ctx, ping)

			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})

		Convey("for failedx2", func() {
			ping := &m.Ping{Type: "http", Count: 2}

			dao.EXPECT().
				ReadPing(ctx, ping.Type).
				Return(ping, nil)

			dao.EXPECT().
				UpdatePing(ctx, ping).
				Return(errt)

			got, err := svc.HandPing(ctx, ping)

			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})
	})
}
