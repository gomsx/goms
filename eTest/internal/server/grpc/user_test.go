package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/aivuca/goms/eTest/api"
	m "github.com/aivuca/goms/eTest/internal/model"
	e "github.com/aivuca/goms/eTest/internal/pkg/err"
	"github.com/aivuca/goms/eTest/internal/service/mock"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := ctxCarryRqid(context.Background())
	errt := errors.New("error")

	Convey("TestCreateUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctx, user).
			Return(nil)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		uidt, err := srv.CreateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
		So(uidt.Uid, ShouldEqual, user.Uid)
	})

	Convey("TestCreateUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		user.Sex = m.GetSexBad()
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.CreateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Sex"])
	})

	Convey("TestCreateUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctx, user).
			Return(errt)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.CreateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := ctxCarryRqid(context.Background())
	errt := errors.New("error")

	Convey("TestReadUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctx, user.Uid).
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

	Convey("TestReadUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		srv.ReadUser(ctx, uidt)
		_, err := srv.ReadUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
	})

	Convey("TestReadUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctx, user.Uid).
			Return(user, errt)
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.ReadUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := ctxCarryRqid(context.Background())
	errt := errors.New("error")

	Convey("TestUpdateUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctx, user).
			Return(nil)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestUpdateUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
	})

	Convey("TestUpdateUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctx, user).
			Return(errt)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := ctxCarryRqid(context.Background())
	errt := errors.New("error")

	Convey("TestDeleteUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(ctx, user.Uid).
			Return(nil)

		//构建 req
		usert := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestDeleteUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
	})

	Convey("TestDeleteUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(ctx, user.Uid).
			Return(errt)
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}
