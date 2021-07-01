package grpc

import (
	"context"

	api "github.com/gomsx/goms/eApi/api/v1"
	m "github.com/gomsx/goms/eApi/internal/model"
	e "github.com/gomsx/goms/pkg/err"
	"github.com/gomsx/goms/pkg/id"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

// setUserReplyMate set reply mate data to user.
func setUserReplyMate(r *api.UserReply, ecode int64, err error) {
	r.Code = ecode
	if err != nil {
		r.Msg = err.Error()
	}
	r.Msg = "ok"
}

// CreateUser create user.
func (s *Server) CreateUser(ctx context.Context, in *api.UserReq) (*api.UserReply, error) {
	// 获取参数
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	// 创建数据
	log.Infof("start to create user, arg: %v", u.String())
	user := &m.User{}
	user.Uid = id.GenUid()
	user.Name = u.GetName()
	user.Sex = u.GetSex()

	// 检验数据
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return res, err
	}
	log.Infof("succeeded to create data, user: %v", *user)

	// 使用数据
	if err := svc.CreateUser(ctx, user); err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Infof("failed to create user, error: %v", err)
		return res, e.ErrInternalError
	}

	// 返回结果
	res.Data.Uid = user.Uid
	setUserReplyMate(res, e.StatusOK, nil)
	log.Infof("succeeded to create user, user: %v", *user)
	return res, nil
}

// ReadUser read user.
func (s *Server) ReadUser(ctx context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data
	log.Infof("start to read user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, err
	}
	log.Infof("succeeded to create data, uid: %v", user.Uid)

	user, err := svc.ReadUser(ctx, user.Uid)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Infof("failed to read user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Data.Uid = user.Uid
	res.Data.Name = user.Name
	res.Data.Sex = user.Sex
	setUserReplyMate(res, e.StatusOK, nil)
	log.Infof("succeeded to read user, user: %v", *user)
	return res, nil
}

// UpdateUser update user.
func (s *Server) UpdateUser(ctx context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data
	log.Infof("start to update user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return res, err
	}
	log.Infof("succeeded to create data, user: %v", *user)

	err := svc.UpdateUser(ctx, user)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Infof("failed to update user, error: %v", err)
		return res, e.ErrInternalError
	}
	setUserReplyMate(res, e.StatusOK, nil)
	log.Infof("succeeded to update user, user: %v", *user)
	return res, nil
}

// DeleteUser delete user.
func (s *Server) DeleteUser(ctx context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data
	log.Infof("start to delete user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, err
	}
	log.Infof("succeeded to create data, uid: %v", user.Uid)

	err := svc.DeleteUser(ctx, user.Uid)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Infof("failed to delete user, error: %v", err)
		return res, e.ErrInternalError
	}
	setUserReplyMate(res, e.StatusOK, err)
	log.Infof("succeeded to delete user, user: %v", *user)
	return res, nil
}

// MapValidateError map validate error.
func MapValidateError(err error) (int64, error) {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	return e.UserEcodeMap[field], e.UserErrMap[field]
}
