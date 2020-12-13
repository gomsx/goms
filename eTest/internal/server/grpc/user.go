package grpc

import (
	"context"

	"github.com/fuwensun/goms/eTest/api"
	m "github.com/fuwensun/goms/eTest/internal/model"
	e "github.com/fuwensun/goms/eTest/internal/pkg/err"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

var empty = &api.Empty{}

// handValidateError hand validate error.
func handValidateError(err error) error {
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		value := ev.Value()
		log.Debug().
			Msgf("arg validate: %v == %v,so error: %v",
				field, value, e.UserErrMap[field])
		return e.UserErrMap[field]
	}
	return nil
}

// CreateUser create user.
func (s *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := s.svc
	res := &api.UidT{}
	// 记录参数
	log.Ctx(c).Info().
		Msgf("start to create user, arg: {%v}", u)

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		// 记录异常
		log.Ctx(c).Info().
			Msgf("failed to validate data, user: %v, error: %v", *user, err)
		return res, handValidateError(err)
	}
	// 记录中间结果
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		// 记录异常
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("failed to create user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	// 记录返回结果
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create user, user = %v", *user)
	return res, nil
}

// ReadUser read user.
func (s *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := s.svc
	res := &api.UserT{}
	log.Ctx(c).Info().
		Msgf("start to read user, arg: {%v}", uid)

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Ctx(c).Info().
			Msgf("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, handValidateError(err)
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		log.Ctx(c).Info().
			Int64("user_id", res.Uid).
			Msgf("failed to read user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	res.Name = user.Name
	res.Sex = user.Sex
	log.Ctx(c).Info().
		Int64("user_id", res.Uid).
		Msgf("succ to read user, user = %v", *user)
	return res, nil
}

// UpdateUser update user.
func (s *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := s.svc
	log.Ctx(c).Info().
		Msgf("start to update user, arg: {%v}", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Ctx(c).Info().
			Msgf("failed to validate data, user: %v, error: %v", *user, err)
		return empty, handValidateError(err)
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("failed to update user, error: %v", err)
		return empty, e.ErrInternalError
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to update user, user = %v", *user)
	return empty, nil
}

// DeleteUser delete user.
func (s *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := s.svc
	log.Ctx(c).Info().
		Msgf("start to delete user, arg: {%v}", uid)

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Ctx(c).Info().
			Msgf("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return empty, handValidateError(err)
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		log.Ctx(c).Info().
			Int64("user_id", user.Uid).
			Msgf("failed to delete user, error: %v", err)
		return empty, e.ErrInternalError
	}
	log.Ctx(c).Info().
		Int64("user_id", user.Uid).
		Msgf("succ to delete user, user = %v", *user)
	return empty, nil
}
