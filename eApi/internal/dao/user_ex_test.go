package dao

import (
	"fmt"
	"reflect"
	"testing"

	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	if isCiEnvDockers() {
		cpstub := gostub.Stub(&cfgpath, "testdata/teardocker/configs")
		defer cpstub.Reset()
	}
	fmt.Printf("==> cfgpath=%v\n", cfgpath)

	// New dao
	dao, clean, err := new(cfgpath)
	if err != nil {
		panic(err)
	}

	Convey("Test dao crud user", t, func() {
		user := m.GetUser()

		err := dao.CreateUser(ctx, user)
		So(err, ShouldBeNil)

		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		user.Name = "bar"
		err = dao.UpdateUser(ctx, user)
		So(err, ShouldBeNil)

		err = dao.DeleteUser(ctx, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test dao crud user db", t, func() {
		user := m.GetUser()

		err := dao.createUserDB(ctx, user)
		So(err, ShouldBeNil)

		got, err := dao.readUserDB(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		user.Name = "bar"
		err = dao.updateUserDB(ctx, user)
		So(err, ShouldBeNil)

		err = dao.deleteUserDB(ctx, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test dao crud user cc", t, func() {
		user := m.GetUser()

		err := dao.setUserCC(ctx, user)
		So(err, ShouldBeNil)

		exist, err := dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)

		got, err := dao.getUserCC(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		err = dao.delUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)

		exist, err = dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)
	})

	Convey("Test dao read user Cache-aside", t, func() {
		user := m.GetUser()

		//create
		err := dao.CreateUser(ctx, user)
		So(err, ShouldBeNil)

		//cache 空
		exist, err := dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)

		//read
		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		//cache 回种
		exist, err = dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)

		//delete
		err = dao.DeleteUser(ctx, user.Uid)
		So(err, ShouldBeNil)

		//cache 失效
		exist, err = dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)
	})

	Convey("Test dao read user Cache-aside", t, func() {
		user := m.GetUser()

		err := dao.CreateUser(ctx, user)
		So(err, ShouldBeNil)

		//cache 空
		exist, err := dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeFalse)

		//read
		got, err := dao.ReadUser(ctx, user.Uid)
		So(reflect.DeepEqual(got, user), ShouldEqual, true)
		So(err, ShouldBeNil)

		//cache 回种
		exist, err = dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)

		//update
		user.Name = "bar"
		err = dao.updateUserDB(ctx, user)
		So(err, ShouldBeNil)

		//cache 回种
		exist, err = dao.existUserCC(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(exist, ShouldBeTrue)
	})

	// 清理
	clean()
}
