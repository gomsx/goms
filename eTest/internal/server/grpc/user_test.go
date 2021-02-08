package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/fuwensun/goms/eTest/api"
	m "github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service/mock"
	e "github.com/fuwensun/goms/pkg/err"
	ms "github.com/fuwensun/goms/pkg/misc"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var errx = errors.New("test error")
var ctxb = context.Background()

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	Convey("TestCreateUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctxb, user).
			Return(nil)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		uidt, err := srv.CreateUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, nil)
		So(uidt.Uid, ShouldEqual, user.Uid)
	})

	Convey("TestCreateUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		user.Sex = ms.GetSexBad()
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.CreateUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Sex"])
	})

	Convey("TestCreateUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctxb, user).
			Return(errx)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.CreateUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	Convey("TestReadUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctxb, user.Uid).
			Return(user, nil)
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		usert, err := srv.ReadUser(ctxb, uidt)
		//断言
		So(err, ShouldEqual, nil)
		So(usert.Uid, ShouldEqual, user.Uid)
		So(usert.Name, ShouldEqual, user.Name)
		So(usert.Sex, ShouldEqual, user.Sex)
	})

	Convey("TestReadUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		srv.ReadUser(ctxb, uidt)
		_, err := srv.ReadUser(ctxb, uidt)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
	})

	Convey("TestReadUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctxb, user.Uid).
			Return(user, errx)
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.ReadUser(ctxb, uidt)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	Convey("TestUpdateUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctxb, user).
			Return(nil)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestUpdateUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
	})

	Convey("TestUpdateUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctxb, user).
			Return(errx)
		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	Convey("TestDeleteUser should StatusOk", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(ctxb, user.Uid).
			Return(nil)
		//构建 req
		usert := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctxb, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestDeleteUser should StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctxb, uidt)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
	})

	Convey("TestDeleteUser should ErrInternalError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(ctxb, user.Uid).
			Return(errx)
		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctxb, uidt)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}
