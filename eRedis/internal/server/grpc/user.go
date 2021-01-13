package grpc

import (
	"context"

	"github.com/fuwensun/goms/eRedis/api"
	m "github.com/fuwensun/goms/eRedis/internal/model"
	e "github.com/fuwensun/goms/pkg/err"
	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/go-playground/validator"
)

var empty = &api.Empty{}

// CreateUser create user.
func (s *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := s.svc
	res := &api.UidT{}

	user := &m.User{}
	user.Uid = ms.GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return res, ms.MapValidateErrorX(err)
	}

	if err := svc.CreateUser(c, user); err != nil {
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	return res, nil
}

// ReadUser read user.
func (s *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := s.svc
	res := &api.UserT{}

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return res, ms.MapValidateErrorX(err)
	}

	user, err := svc.ReadUser(c, uid.Uid)
	if err != nil {
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid
	res.Name = user.Name
	res.Sex = user.Sex
	return res, nil
}

// UpdateUser update user.
func (s *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := s.svc

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return empty, ms.MapValidateErrorX(err)
	}

	err := svc.UpdateUser(c, user)
	if err != nil {
		return empty, e.ErrInternalError
	}
	return empty, nil
}

// DeleteUser delete user.
func (s *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := s.svc

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return empty, ms.MapValidateErrorX(err)
	}

	err := svc.DeleteUser(c, uid.Uid)
	if err != nil {
		return empty, e.ErrInternalError
	}
	return empty, nil
}
