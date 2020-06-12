package service

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eApi/internal/dao/mock"
	. "github.com/fuwensun/goms/eApi/internal/model"

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
			p := &{
				Type:"http",
				Count:2,
			}
			wang := &{
				Type:"http",
				Count:3,
			}
			daom.EXPECT().
				ReadPing(gomock.Any(), p.Type).
				Return(p, nil)

			daom.EXPECT().
				UpdatePing(gomock.Any(), p).
				Return(nil)

			got, err := svc.HandPing(context.Background())
			So(reflect.DeepEqual(got,want),ShouldEqual,true)
			So(err, ShouldBeNil)
		})
	})
}