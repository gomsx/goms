package grpc

import (
	"context"

	"github.com/fuwensun/goms/eRedis/api"
	. "github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/go-playground/validator"
)

var empty = &api.Empty{}

func handValidateError(err error) error {
	for _, ev := range err.(validator.ValidationErrors) {
		return UserErrMap[ev.Namespace()]
	}
	return nil
}

// createUser
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc
	res := &api.UidT{}

	user := &User{}
	user.Uid = GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return res, handValidateError(err)
	}

	if err := svc.CreateUser(c, user); err != nil {
		return res, ErrInternalError
	}
	res.Uid = user.Uid

	return res, nil
}

// readUser
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc
	res := &api.UserT{}

	user := &User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return res, handValidateError(err)
	}

	u, err := svc.ReadUser(c, uid.Uid)
	if err != nil {
		return res, ErrInternalError
	}

	res.Uid = u.Uid
	res.Name = u.Name
	res.Sex = u.Sex

	return res, nil
}

// updateUser
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc

	user := &User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return empty, handValidateError(err)
	}

	err := svc.UpdateUser(c, user)
	if err != nil {
		return empty, ErrInternalError
	}

	return empty, nil
}

// deleteUser
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc

	user := &User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return empty, handValidateError(err)
	}

	err := svc.DeleteUser(c, uid.Uid)
	if err != nil {
		return empty, ErrInternalError
	}

	return empty, nil
}
