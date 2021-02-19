package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/aivuca/goms/eTest/api"
	m "github.com/aivuca/goms/eTest/internal/model"
	"github.com/aivuca/goms/eTest/internal/service/mock"
	e "github.com/aivuca/goms/pkg/err"
	ms "github.com/aivuca/goms/pkg/misc"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var errx = errors.New("error xxx")
var ctxb = context.Background()
var ctxa = gomock.Any()

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}

	Convey("Create user with correct user data", t, func() {
		//mock
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctxa, user).
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

	Convey("Create user with incorrect user data", t, func() {
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

	Convey("Create user when InternalServerError", t, func() {
		//mock
		user := m.GetUser()
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctxa, user).
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

	Convey("Read user with correct user data", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctxa, user.Uid).
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

	Convey("Read user with incorrect user data", t, func() {
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

	Convey("Read user when InternalServerError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctxa, user.Uid).
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

	Convey("Update user with correct user data", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctxa, user).
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

	Convey("Update user with incorrect user data", t, func() {
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

	Convey("Update user when InternalServerError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctxa, user).
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

	Convey("Delete user with correct user data", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(ctxa, user.Uid).
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

	Convey("Delete user with incorrect user data", t, func() {
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

	Convey("DeleteUser should InternalServerError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			DeleteUser(ctxa, user.Uid).
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
