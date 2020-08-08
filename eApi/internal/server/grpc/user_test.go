package grpc

import (
	"context"
	"testing"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	"github.com/aivuca/goms/eApi/internal/service/mock"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var ctx = context.Background()

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestCreateUser should succ", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(nil)

		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUserX(ctx, req)

		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, 200)
		So(res.Data.Uid, ShouldEqual, user.Uid)

	})

	Convey("TestCreateUser should failed", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(e.ErrInternalError)

		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUserX(ctx, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, 500)
		So(res.Data.Uid, ShouldEqual, 0) //todo

	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestReadUser should succ", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, nil)

		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUserX(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, 200)
		So(res.Data.Uid, ShouldEqual, user.Uid)
		So(res.Data.Name, ShouldEqual, user.Name)
		So(res.Data.Sex, ShouldEqual, user.Sex)
	})

	Convey("TestReadUser should failed", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, e.ErrInternalError)

		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUserX(ctx, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, 500)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestUpdateUser should succ", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(nil)

		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		_, err := srv.UpdateUserX(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestUpdateUser should failed", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(e.ErrInternalError)

		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		_, err := srv.UpdateUserX(ctx, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestDeleteUser should succ", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(gomock.Any(), user.Uid).
			Return(nil)

		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		_, err := srv.DeleteUserX(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestDeleteUser should failed", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(gomock.Any(), user.Uid).
			Return(e.ErrInternalError)

		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		_, err := srv.DeleteUserX(ctx, req)

		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}
