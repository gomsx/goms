package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/aivuca/goms/eApi/internal/dao/mock"
	. "github.com/aivuca/goms/eApi/internal/model"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

//
func TestHandPing(t *testing.T) {
	Convey("TestHandPing", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		Convey("for succ", func() {
			p := &Ping{
				Type:  "http",
				Count: 2,
			}
			want := &Ping{
				Type:  "http",
				Count: 3,
			}
			daom.EXPECT().
				ReadPing(gomock.Any(), p.Type).
				Return(p, nil)

			daom.EXPECT().
				UpdatePing(gomock.Any(), p).
				Return(nil)

			got, err := svc.HandPing(context.Background(), p)
			So(reflect.DeepEqual(got, want), ShouldEqual, true)
			So(err, ShouldBeNil)
		})
	})
}

