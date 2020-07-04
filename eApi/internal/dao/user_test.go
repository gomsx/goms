package dao_test

//dao_test 外部测试包，包名是 dao_test,不是 dao
import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/fuwensun/goms/eApi/internal/dao"
	. "github.com/fuwensun/goms/eApi/internal/model"
)

var cfgpath = "testdata/configs"
var ctx = context.Background()

func TestUser(t *testing.T) {
	// 读取配置
	if CI_ENV_NO_DOCKER == "" {
		cpstub := gostub.Stub(&cfgpath, "testdata/teardocker/configs")
		defer cpstub.Reset()
	}
	fmt.Printf("==> cfgpath=%v\n", cfgpath)

	// New dao
	dao, clean, err := New(cfgpath)
	if err != nil {
		panic(err)
	}

	Convey("Test dao crud user", t, func() {

		user := &User{Name: "foo", Sex: 0}
		user.Uid = GetUid()

		err := dao.CreateUser(ctx, user)
		So(err, ShouldBeNil)

		user.Name = "bar"
		err = dao.UpdateUser(ctx, user)
		So(err, ShouldBeNil)

		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		err = dao.DeleteUser(ctx, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test dao crud user db", t, func() {

		user := &User{Name: "foo", Sex: 0}
		user.Uid = GetUid()

		err := dao.CreateUserDB(ctx, user)
		So(err, ShouldBeNil)

		user.Name = "bar"
		err = dao.UpdateUserDB(ctx, user)
		So(err, ShouldBeNil)

		got, err := dao.ReadUserDB(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		err = dao.DeleteUserDB(ctx, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test dao crud user cc", t, func() {

		user := &User{Name: "foo", Sex: 0}
		user.Uid = GetUid()

		err := dao.SetUserCC(ctx, user)
		So(err, ShouldBeNil)

		exist, err := dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)

		got, err := dao.GetUserCC(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		err = dao.DelUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)

		exist, err = dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)
	})

	Convey("Test dao read user Cache-aside", t, func() {

		user := &User{Name: "foo", Sex: 0}
		user.Uid = GetUid()

		//create
		err := dao.CreateUser(ctx, user)
		So(err, ShouldBeNil)

		//cache 空
		exist, err := dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)

		//read
		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		//cache 回种
		exist, err = dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)

		//delete
		err = dao.DeleteUser(ctx, user.Uid)
		So(err, ShouldBeNil)

		//cache 失效
		exist, err = dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)
	})

	Convey("Test dao read user Cache-aside", t, func() {

		user := &User{Name: "foo", Sex: 0}
		user.Uid = GetUid()

		err := dao.CreateUser(ctx, user)
		So(err, ShouldBeNil)

		//cache 空
		exist, err := dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)

		//read
		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		//cache 回种
		exist, err = dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)

		//update
		user.Name = "bar"
		err = dao.UpdateUserDB(ctx, user)
		So(err, ShouldBeNil)

		//cache 回种
		exist, err = dao.ExistUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)
	})

	// 清理
	clean()
}

