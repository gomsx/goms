package grpc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	api "github.com/fuwensun/goms/eApi/api/v1"
	. "github.com/fuwensun/goms/eApi/internal/model"
	"github.com/fuwensun/goms/eApi/internal/service/mock"
)

var ctx = context.Background()

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestCreateUser should succ", t, func() {
		//mock
		user := &User{
			// Uid:  123,
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(nil)

		//构建 req
		usert := &api.UserT{
			// Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		uidt, err := srv.CreateUser(ctx, usert)

		//断言
		So(uidt.Uid, ShouldEqual, 0)
		So(err, ShouldEqual, nil)
	})

	Convey("TestCreateUser should failed", t, func() {
		//mock
		user := &User{
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(ErrInternalError)

		//构建 req
		usert := &api.UserT{
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		uidt, err := srv.CreateUser(ctx, usert)
		//断言
		So(uidt.Uid, ShouldEqual, 0)
		So(err, ShouldEqual, ErrInternalError)
	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestReadUser should succ", t, func() {
		//mock
		user := &User{
			Uid:  123,
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, nil)

		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		usert, err := srv.ReadUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, nil)
		So(usert.Uid, ShouldEqual, user.Uid)
		So(usert.Name, ShouldEqual, user.Name)
		So(usert.Sex, ShouldEqual, user.Sex)
	})

	Convey("TestReadUser should failed", t, func() {
		//mock
		user := &User{
			Uid:  123,
			Name: "xxx",
			Sex:  0,
		}

		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, ErrInternalError)

		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.ReadUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, ErrInternalError)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestUpdateUser should succ", t, func() {
		//mock
		user := &User{
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(nil)

		//构建 req
		usert := &api.UserT{
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestUpdateUser should failed", t, func() {
		//mock
		user := &User{
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(ErrInternalError)

		//构建 req
		usert := &api.UserT{
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, ErrInternalError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestDeleteUser should succ", t, func() {
		//mock
		user := &User{
			Uid:  123,
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			DeleteUser(gomock.Any(), user.Uid).
			Return(nil)

		//构建 req
		var ctx = context.Background()
		usert := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestDeleteUser should failed", t, func() {
		//mock
		user := &User{
			Uid:  123,
			Name: "xxx",
			Sex:  0,
		}
		svcm.EXPECT().
			DeleteUser(gomock.Any(), user.Uid).
			Return(ErrInternalError)

		//构建 req
		var ctx = context.Background()
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, uidt)

		//断言
		So(err, ShouldEqual, ErrInternalError)
	})
}

