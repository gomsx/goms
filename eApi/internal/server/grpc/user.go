package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

// handValidataError.
func handValidataError(c context.Context, err error) (int64, error) {
	// for _, ev := range err.(validator.ValidationErrors) {...}//todo
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		log.Ctx(c).Debug().
			Msgf("arg validate error: %v==%v", ev.StructField(), ev.Value())
		return e.UserEcodeMap[field], e.UserErrMap[field]
	}
	return 0, nil
}

//
func setUserReplyMate(r *api.UserReply, ecode int64, err error) {
	r.Code = ecode
	if err != nil {
		r.Msg = err.Error()
	}
	r.Msg = "ok"
}

// CreateUser create user.
func (srv *Server) CreateUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	// 获取参数
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	// 创建数据
	log.Ctx(c).Info().
		Msgf("start to create user, arg: %v", u.String())

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = u.GetName()
	user.Sex = u.GetSex()

	// 检验数据
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := handValidataError(c, err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return res, err
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	// 使用数据
	if err := svc.CreateUser(c, user); err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to create user, data: %v, error: %v", *user, err)
		return res, e.ErrInternalError
	}
	res.Data.Uid = user.Uid
	setUserReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create user, user = %v", *user)
	return res, nil
}

// ReadUser read user.
func (srv *Server) ReadUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Ctx(c).Info().
		Msgf("start to read user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := handValidataError(c, err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return res, err
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return res, e.ErrInternalError
	}
	res.Data.Uid = user.Uid
	res.Data.Name = user.Name
	res.Data.Sex = user.Sex
	setUserReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to read user, user = %v", *user)
	return res, nil
}

// UpdateUser update user.
func (srv *Server) UpdateUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Ctx(c).Info().
		Msgf("start to update user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := handValidataError(c, err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return res, err
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to update user, data: %v, error: %v", *user, err)
		return res, e.ErrInternalError
	}
	setUserReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to update user, user = %v", *user)
	return res, nil
}

// DeleteUser delete user.
func (srv *Server) DeleteUser(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Ctx(c).Info().
		Msgf("start to delete user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := handValidataError(c, err)
		setUserReplyMate(res, ecode, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return res, err
	}

	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		setUserReplyMate(res, e.StatusInternalServerError, err)
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return res, e.ErrInternalError
	}
	setUserReplyMate(res, e.StatusOK, err)
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to read user, user = %v", *user)
	return res, nil
}
