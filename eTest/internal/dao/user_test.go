package dao_test

//dao_test 外部测试包，包名是 dao_test,不是 dao
import (
	"context"
	"fmt"
	"reflect"
	"testing"

	. "github.com/fuwensun/goms/eTest/internal/dao"
	. "github.com/fuwensun/goms/eTest/internal/model"
	svc "github.com/fuwensun/goms/eTest/internal/service"
	. "github.com/smartystreets/goconvey/convey"
)

var cfgpath = "testdata/configs"
var ctx = context.Background()

func TestDao(t *testing.T) {
	fmt.Printf("==> cfgpath=%v\n", cfgpath)
	dao, clean, err := New(cfgpath)
	if err != nil {
		panic(err)
	}
	svc.InitUidGenerator()
	Convey("Test dao curd user", t, func() {

		user := User{Name: "x1", Sex: 0}
		user.Uid = svc.GetUid()
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
