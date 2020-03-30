package dao

import (
	"context"
	"math/rand"
	"reflect"
	"testing"
	"time"

	. "github.com/fuwensun/goms/eTest/internal/model"
	. "github.com/smartystreets/goconvey/convey"
)

var ctx = context.Background()

func TestDao(t *testing.T) {
	dao, clean, err := New("testdata/tearA/configs")
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	Convey("Test dao", t, func() {

		user := User{Name: "x1", Sex: 0}
		user.Uid = rand.Int63n(0x0FFF_FFFF_FFFF_FFFF)
		err := dao.CreateUser(ctx, &user)
		So(err, ShouldBeNil)

		user.Name = "x2"
		err = dao.UpdateUser(ctx, &user)
		So(err, ShouldBeNil)

		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		err = dao.DeleteUser(ctx, user.Uid)
		So(err, ShouldBeNil)
	})
	clean()
}
