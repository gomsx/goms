package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/pkg/err"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
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
func (s *Server) CreateUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	// 获取参数
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	// 创建数据
	log.Ctx(c).Info().Msgf("start to create user, arg: %v", u.String())
	user := &m.User{}
	user.Uid = ms.GetUid()
	user.Name = u.GetName()
	user.Sex = u.GetSex()

	// 检验数据
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := ms.MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().Msgf("failed to validate data, user: %v, error: %v", *user, err)
		return res, err
	}
	log.Ctx(c).Info().Msgf("succ to create data, user: %v", *user)

	// 使用数据
	c = ms.CarryCtxUserId(c, user.Uid)
	if err := svc.CreateUser(c, user); err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().Msgf("failed to create user, error: %v", err)
		return res, e.ErrInternalError
	}

	// 返回结果
	res.Data.Uid = user.Uid
	setUserReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Info().Msgf("succ to create user, user: %v", *user)
	return res, nil
}

// ReadUser read user.
func (s *Server) ReadUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data
	log.Ctx(c).Info().Msgf("start to read user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := ms.MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().Msgf("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, err
	}
	log.Ctx(c).Info().Msgf("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().Msgf("failed to read user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Data.Uid = user.Uid
	res.Data.Name = user.Name
	res.Data.Sex = user.Sex
	setUserReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Info().Msgf("succ to read user, user: %v", *user)
	return res, nil
}

// UpdateUser update user.
func (s *Server) UpdateUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data
	log.Ctx(c).Info().Msgf("start to update user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := ms.MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().Msgf("failed to validate data, user: %v, error: %v", *user, err)
		return res, err
	}
	log.Ctx(c).Info().Msgf("succ to create data, user: %v", *user)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.UpdateUser(c, user)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().Msgf("failed to update user, error: %v", err)
		return res, e.ErrInternalError
	}
	setUserReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Info().Msgf("succ to update user, user: %v", *user)
	return res, nil
}

// DeleteUser delete user.
func (s *Server) DeleteUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := s.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data
	log.Ctx(c).Info().Msgf("start to delete user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := ms.MapValidateError(err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().Msgf("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, err
	}
	log.Ctx(c).Info().Msgf("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().Msgf("failed to delete user, error: %v", err)
		return res, e.ErrInternalError
	}
	setUserReplyMate(res, e.StatusOK, err)
	log.Ctx(c).Info().Msgf("succ to delete user, user: %v", *user)
	return res, nil
}
