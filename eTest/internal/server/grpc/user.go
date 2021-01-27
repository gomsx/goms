package grpc

import (
	"context"

	"github.com/aivuca/goms/eTest/api"
	m "github.com/aivuca/goms/eTest/internal/model"
	e "github.com/aivuca/goms/pkg/err"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

var empty = &api.Empty{}

// CreateUser create user.
func (s *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := s.svc
	res := &api.UidT{}
	// 记录参数
	log.Infof("start to create user, arg: {%v}", u)

	user := &m.User{}
	user.Uid = ms.GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		// 记录异常
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return res, ms.MapValidateErrorX(err)
	}
	// 记录中间结果
	log.Infof("succ to create data, user: %v", *user)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.CreateUser(c, user)
	if err != nil {
		// 记录异常
		log.Infof("failed to create user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	// 记录返回结果
	log.Infof("succ to create user, user: %v", *user)
	return res, nil
}

// ReadUser read user.
func (s *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := s.svc
	res := &api.UserT{}
	log.Infof("start to read user, arg: {%v}", uid)

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return res, ms.MapValidateErrorX(err)
	}
	log.Infof("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		log.Infof("failed to read user, error: %v", err)
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	res.Name = user.Name
	res.Sex = user.Sex
	log.Infof("succ to read user, user: %v", *user)
	return res, nil
}

// UpdateUser update user.
func (s *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := s.svc
	log.Infof("start to update user, arg: {%v}", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return empty, ms.MapValidateErrorX(err)
	}
	log.Infof("succ to create data, user: %v", *user)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.UpdateUser(c, user)
	if err != nil {
		log.Infof("failed to update user, error: %v", err)
		return empty, e.ErrInternalError
	}
	log.Infof("succ to update user, user: %v", *user)
	return empty, nil
}

// DeleteUser delete user.
func (s *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := s.svc
	log.Infof("start to delete user, arg: {%v}", uid)

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return empty, ms.MapValidateErrorX(err)
	}
	log.Infof("succ to create data, uid: %v", user.Uid)

	c = ms.CarryCtxUserId(c, user.Uid)
	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		log.Infof("failed to delete user, error: %v", err)
		return empty, e.ErrInternalError
	}
	log.Infof("succ to delete user, user: %v", *user)
	return empty, nil
}
