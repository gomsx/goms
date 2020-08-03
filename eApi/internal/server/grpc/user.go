package grpc

import (
	"context"

	api "github.com/fuwensun/goms/eApi/api/v1"
	. "github.com/fuwensun/goms/eApi/internal/model"
	"github.com/fuwensun/goms/eApi/internal/pkg/reqid"

	"github.com/go-playground/validator"
)

var empty = &api.Empty{}

func handValidateError(c context.Context, err error) error {
	for _, ev := range err.(validator.ValidationErrors) {
		log.Debug().
			Int64("request_id", reqid.GetIdMust(c)).
			Msgf("%v err => %v", ev.StructField(), ev.Value())
		return UserErrMap[ev.Namespace()]
	}
	return nil
}

// createUser
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc
	res := &api.UidT{}

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("start to create user,arg: %v", u)

	user := &User{}
	user.Uid = GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return res, handValidateError(c, err)
	}

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		log.Info().
			Int64("request_id", reqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to create user")
		return res, ErrInternalError
	}
	res.Uid = user.Uid

	log.Info().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to create user")
	return res, nil
}

// readUser
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc
	res := &api.UserT{}

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msg("start to read user")

	user := &User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return res, handValidateError(c, err)
	}

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("succ to get user uid, uid = %v", uid)

	u, err := svc.ReadUser(c, uid.Uid)
	if err != nil {
		log.Info().
			Int64("request_id", reqid.GetIdMust(c)).
			Int64("user_id", res.Uid).
			Msg("failed to read user")
		return res, ErrInternalError
	}

	res.Uid = u.Uid
	res.Name = u.Name
	res.Sex = u.Sex

	log.Info().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", res.Uid).
		Msg("succ to read user")
	return res, nil
}

// updateUser
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("start to update user, arg: %v", u)

	user := &User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return empty, handValidateError(c, err)
	}

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		log.Info().
			Int64("request_id", reqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to update user")
		return empty, ErrInternalError
	}
	log.Info().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to update user")
	return empty, nil
}

// deleteUser
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msg("start to delete user")

	user := &User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return empty, handValidateError(c, err)
	}

	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Msgf("succ to get user uid, uid = %v", uid)

	err := svc.DeleteUser(c, uid.Uid)
	if err != nil {
		log.Info().
			Int64("request_id", reqid.GetIdMust(c)).
			Int64("user_id", uid.Uid).
			Msg("failed to delete user")
		return empty, ErrInternalError
	}

	log.Info().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", uid.Uid).
		Msg("failed to delete user")
	return empty, nil
}
