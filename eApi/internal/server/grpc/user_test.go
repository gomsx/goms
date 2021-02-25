package grpc

import (
	"context"
	"errors"
	"testing"

	api "github.com/fuwensun/goms/eApi/api/v1"
	m "github.com/fuwensun/goms/eApi/internal/model"
	"github.com/fuwensun/goms/eApi/internal/service/mock"
	e "github.com/fuwensun/goms/pkg/err"
	ms "github.com/fuwensun/goms/pkg/misc"

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
		user := m.GetUser()
		//patch
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctxa, user).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Data.Uid, ShouldEqual, user.Uid)
	})

	Convey("Create user with incorrect user data", t, func() {
		user := m.GetUser()
		//patch
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		user.Sex = ms.GetSexBad()
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Sex"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Sex"])
	})

	Convey("Create user when InternalServerError", t, func() {
		user := m.GetUser()
		//patch
		Patch(ms.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctxa, user).
			Return(errx)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, e.StatusInternalServerError)
	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("Read user with correct user data", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			ReadUser(ctxa, user.Uid).
			Return(user, nil)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Data.Uid, ShouldEqual, user.Uid)
		So(res.Data.Name, ShouldEqual, user.Name)
		So(res.Data.Sex, ShouldEqual, user.Sex)
	})

	Convey("Read user with incorrect user data", t, func() {
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("Read user when InternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			ReadUser(ctxa, user.Uid).
			Return(user, errx)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, e.StatusInternalServerError)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("Update user with correct user data", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			UpdateUser(ctxa, user).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
	})

	Convey("Update user with incorrect user data", t, func() {
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("Update user when InternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			UpdateUser(ctxa, user).
			Return(errx)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, e.StatusInternalServerError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("Delete user with correct user data", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			DeleteUser(ctxa, user.Uid).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
	})

	Convey("Delete user with incorrect user data", t, func() {
		user := m.GetUser()
		user.Uid = ms.GetUidBad()
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("Delete user when InternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			DeleteUser(ctxa, user.Uid).
			Return(errx)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctxb, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, e.StatusInternalServerError)
	})
}
