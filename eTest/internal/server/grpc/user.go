package grpc

import (
	"context"

	"github.com/fuwensun/goms/eTest/api"
	m "github.com/fuwensun/goms/eTest/internal/model"
	e "github.com/fuwensun/goms/pkg/err"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

var empty = &api.Empty{}

// CreateUser create user.
func (s *Server) CreateUser(ctx context.Context, u *api.UserT) (*api.UidT, error) {
	svc := s.svc
	res := &api.UidT{}
	// 记录参数
	log.Infof("start to create user, arg: {%v}", u)

	user := &m.User{}
	user.Uid = ms.GenUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		// 记录异常
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return res, MapValidateError(err)
	}
	// 记录中间结果
	log.Infof("succeeded to create data, user: %v", *user)

	err := svc.CreateUser(ctx, user)
	if err != nil {
		// 记录异常
		log.Infof("failed to create user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	// 记录返回结果
	log.Infof("succeeded to create user, user: %v", *user)
	return res, nil
}

// ReadUser read user.
func (s *Server) ReadUser(ctx context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := s.svc
	res := &api.UserT{}
	log.Infof("start to read user, arg: {%v}", uid)

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, MapValidateError(err)
	}
	log.Infof("succeeded to create data, uid: %v", user.Uid)

	user, err := svc.ReadUser(ctx, user.Uid)
	if err != nil {
		log.Infof("failed to read user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	res.Name = user.Name
	res.Sex = user.Sex
	log.Infof("succeeded to read user, user: %v", *user)
	return res, nil
}

// UpdateUser update user.
func (s *Server) UpdateUser(ctx context.Context, u *api.UserT) (*api.Empty, error) {
	svc := s.svc
	log.Infof("start to update user, arg: {%v}", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return empty, MapValidateError(err)
	}
	log.Infof("succeeded to create data, user: %v", *user)

	err := svc.UpdateUser(ctx, user)
	if err != nil {
		log.Infof("failed to update user, error: %v", err)
		return empty, e.ErrInternalError
	}
	log.Infof("succeeded to update user, user: %v", *user)
	return empty, nil
}

// DeleteUser delete user.
func (s *Server) DeleteUser(ctx context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := s.svc
	log.Infof("start to delete user, arg: {%v}", uid)

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return empty, MapValidateError(err)
	}
	log.Infof("succeeded to create data, uid: %v", user.Uid)

	err := svc.DeleteUser(ctx, user.Uid)
	if err != nil {
		log.Infof("failed to delete user, error: %v", err)
		return empty, e.ErrInternalError
	}
	log.Infof("succeeded to delete user, user: %v", *user)
	return empty, nil
}

// MapValidateError map validate error.
func MapValidateError(err error) error {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	return e.UserErrMap[field]
}
