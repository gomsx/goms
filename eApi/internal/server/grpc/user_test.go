package grpc

import (
	"context"
	"errors"
	"testing"

	api "github.com/fuwensun/goms/eApi/api/v1"
	m "github.com/fuwensun/goms/eApi/internal/model"
	e "github.com/fuwensun/goms/eApi/internal/pkg/err"
	"github.com/fuwensun/goms/eApi/internal/service/mock"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var errx = errors.New("test error")
var ctxb = context.Background()
var ctxq = carryCtxRequestId(ctxb)

// var ctxq = gomock.Any() // struct{} 接受任意 ctx

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestCreateUser should StatusOK", t, func() {
		user := m.GetUser()

		//patch
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctxq, user).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctxq, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Data.Uid, ShouldEqual, user.Uid)
	})

	Convey("TestCreateUser should StatusBadRequest", t, func() {
		user := m.GetUser()
		//patch
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		user.Sex = m.GetSexBad()
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctxq, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Sex"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Sex"])
	})

	Convey("TestCreateUser should StatusInternalServerError", t, func() {
		user := m.GetUser()

		//patch
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctxq, user).
			Return(errx)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctxq, req)
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
			ReadUser(ctxq, user.Uid).
			Return(user, nil)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctxq, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
		So(res.Data.Uid, ShouldEqual, user.Uid)
		So(res.Data.Name, ShouldEqual, user.Name)
		So(res.Data.Sex, ShouldEqual, user.Sex)
	})

	Convey("TestReadUser should StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		ctxq := carryCtxRequestId(ctxb)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctxq, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("TestReadUser should StatusInternalServerError", t, func() {
		user := m.GetUser()

		//mock
		svcm.EXPECT().
			ReadUser(ctxq, user.Uid).
			Return(user, errx)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctxq, req)
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
			UpdateUser(ctxq, user).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctxq, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
	})

	Convey("TestUpdateUser should StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		ctxq := carryCtxRequestId(ctxb)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctxq, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("TestUpdateUser should StatusInternalServerError", t, func() {
		user := m.GetUser()

		//mock
		svcm.EXPECT().
			UpdateUser(ctxq, user).
			Return(errx)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctxq, req)
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
			DeleteUser(ctxq, user.Uid).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctxq, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
	})

	Convey("TestDeleteUser should StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctxq, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("TestDeleteUser should ErrInternalError", t, func() {
		user := m.GetUser()

		//mock
		svcm.EXPECT().
			DeleteUser(ctxq, user.Uid).
			Return(errx)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctxq, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, e.StatusInternalServerError)
	})
}
