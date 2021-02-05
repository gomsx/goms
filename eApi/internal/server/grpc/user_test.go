package grpc

import (
	"context"
	"errors"
	"testing"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	"github.com/aivuca/goms/eApi/internal/service/mock"
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

	Convey("TestCreateUser should StatusOK", t, func() {
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

	Convey("TestCreateUser should StatusBadRequest", t, func() {
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

	Convey("TestCreateUser should StatusInternalServerError", t, func() {
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

	Convey("TestReadUser should StatusOK", t, func() {
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

	Convey("TestReadUser should StatusBadRequest", t, func() {
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

	Convey("TestReadUser should StatusInternalServerError", t, func() {
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

	Convey("TestUpdateUser should StatusOK", t, func() {
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

	Convey("TestUpdateUser should StatusBadRequest", t, func() {
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

	Convey("TestUpdateUser should StatusInternalServerError", t, func() {
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

	Convey("TestDeleteUser should StatusOK", t, func() {
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

	Convey("TestDeleteUser should StatusBadRequest", t, func() {
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

	Convey("TestDeleteUser should ErrInternalError", t, func() {
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
