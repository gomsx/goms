package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aivuca/goms/eApi/internal/dao/mock"
	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandPing(t *testing.T) {
	Convey("TestHandPing", t, func() {
		ctx := context.Background()
		//
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		dao := mock.NewMockDao(ctrl)
		svc := service{dao: dao}

		Convey("for succ", func() {
			p := &m.Ping{Type: "http", Count: 2}
			want := &m.Ping{Type: "http", Count: 3}

			dao.EXPECT().
				ReadPing(ctx, p.Type).
				Return(p, nil)

			dao.EXPECT().
				UpdatePing(ctx, p).
				Return(nil)

			got, err := svc.HandPing(ctx, p)

			So(err, ShouldBeNil)
			So(reflect.DeepEqual(got, want), ShouldBeTrue)
		})

		Convey("for failed", func() {
			p := &m.Ping{Type: "http", Count: 2}

			dao.EXPECT().
				ReadPing(ctx, p.Type).
				Return(p, errors.New("xxx"))

			got, err := svc.HandPing(ctx, p)

			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})

		Convey("for failedx2", func() {
			p := &m.Ping{Type: "http", Count: 2}

			dao.EXPECT().
				ReadPing(ctx, p.Type).
				Return(p, nil)

			dao.EXPECT().
				UpdatePing(ctx, p).
				Return(errors.New("xxx"))

			got, err := svc.HandPing(ctx, p)

			So(err, ShouldNotBeNil)
			So(got, ShouldBeNil)
		})
	})
}
