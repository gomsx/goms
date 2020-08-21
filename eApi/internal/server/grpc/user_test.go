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

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := ctxCarryRqid(context.Background())

	Convey("TestCreateUser should StatusOK", t, func() {
		user := m.GetUser()
		//patch
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		//mock
		svcm.EXPECT().
			CreateUser(ctx, user).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctx, req)
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
		//mock
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctx, req)
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
			CreateUser(ctx, user).
			Return(e.ErrInternalError)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.CreateUser(ctx, req)
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
	ctx := ctxCarryRqid(context.Background())

	Convey("TestReadUser should StatusOK", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			ReadUser(ctx, user.Uid).
			Return(user, nil)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctx, req)
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
		//mock
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctx, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("TestReadUser should StatusInternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			ReadUser(ctx, user.Uid).
			Return(user, e.ErrInternalError)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.ReadUser(ctx, req)
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
	ctx := ctxCarryRqid(context.Background())

	Convey("TestUpdateUser should StatusOK", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			UpdateUser(ctx, user).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
	})

	Convey("TestUpdateUser should StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		//mock
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctx, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("TestUpdateUser should StatusInternalServerError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			UpdateUser(ctx, user).
			Return(e.ErrInternalError)
		//构建 req
		data := &api.UserMsg{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.UpdateUser(ctx, req)
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
	ctx := ctxCarryRqid(context.Background())

	Convey("TestDeleteUser should StatusOK", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			DeleteUser(ctx, user.Uid).
			Return(nil)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctx, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Code, ShouldEqual, e.StatusOK)
	})

	Convey("TestDeleteUser should StatusBadRequest", t, func() {
		user := m.GetUser()
		user.Uid = m.GetUidBad()
		//mock
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctx, req)
		//断言
		So(err, ShouldEqual, e.UserErrMap["Uid"])
		So(res.Code, ShouldEqual, e.UserEcodeMap["Uid"])
	})

	Convey("TestDeleteUser should ErrInternalError", t, func() {
		user := m.GetUser()
		//mock
		svcm.EXPECT().
			DeleteUser(ctx, user.Uid).
			Return(e.ErrInternalError)
		//构建 req
		data := &api.UserMsg{
			Uid: user.Uid,
		}
		req := &api.UserReq{Data: data}
		//发起 req
		res, err := srv.DeleteUser(ctx, req)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
		So(res.Code, ShouldEqual, e.StatusInternalServerError)
	})
}
