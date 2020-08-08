package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"

	"github.com/go-playground/validator"
)

//
var empty = &api.Empty{}

// handValidataError.
func handValidataError(c context.Context, err error) error {
	// for _, ev := range err.(validator.ValidationErrors) {...}//todo
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		log.Debug().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("arg validate error: %v==%v", ev.StructField(), ev.Value())
		return e.UserErrMap[field]
	}
	return nil
}

// handValidataError.
func handValidataErrorX(c context.Context, err error) (int64, error) {
	// for _, ev := range err.(validator.ValidationErrors) {...}//todo
	if ev := err.(validator.ValidationErrors)[0]; ev != nil {
		field := ev.StructField()
		log.Debug().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("arg validate error: %v==%v", ev.StructField(), ev.Value())
		return e.UserEcodeMap[field], e.UserErrMap[field]
	}
	return 0, nil
}

func setReplyMate(r *api.UserReply, ecode int64, err error) {
	r.Code = ecode
	if err != nil {
		r.Msg = err.Error()
	}
}

// CreateUser create user.
func (srv *Server) CreateUserX(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("start to create user,arg: %v", in)

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := handValidataErrorX(c, err)
		setReplyMate(res, ecode, err)
		return res, err
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		setReplyMate(res, 500, e.ErrInternalError)
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to create user")
		return res, e.ErrInternalError
	}

	res.Data.Uid = user.Uid
	setReplyMate(res, 200, nil)
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to create user")

	return res, nil
}

// ReadUser read user.
func (srv *Server) ReadUserX(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to read user")

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := handValidataErrorX(c, err)
		setReplyMate(res, ecode, err)
		return res, err
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user data, uid = %v", user.Uid)

	user, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		setReplyMate(res, 500, e.ErrInternalError)
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to read user")
		return res, e.ErrInternalError
	}

	res.Data.Uid = user.Uid
	res.Data.Name = user.Name
	res.Data.Sex = user.Sex
	setReplyMate(res, 200, nil)
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to read user")
	return res, nil
}

// UpdateUser update user.
func (srv *Server) UpdateUserX(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("start to update user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ecode, err := handValidataErrorX(c, err)
		setReplyMate(res, ecode, err)
		return res, err
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		setReplyMate(res, 500, e.ErrInternalError)
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to update user")
		return res, e.ErrInternalError
	}
	setReplyMate(res, 200, nil)
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to update user")
	return res, nil
}

// DeleteUser delete user.
func (srv *Server) DeleteUserX(c context.Context, in *api.UserReq) (*api.UserReply, error) {
	svc := srv.svc
	res := &api.UserReply{Data: &api.UserMsg{}}
	u := in.Data

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to delete user")

	user := &m.User{}
	user.Uid = u.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ecode, err := handValidataErrorX(c, err)
		setReplyMate(res, ecode, err)
		return res, err
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user uid, uid = %v", user.Uid)

	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		setReplyMate(res, 500, err)
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to delete user")
		return res, e.ErrInternalError
	}

	setReplyMate(res, 200, err)
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("failed to delete user")
	return res, nil
}

// CreateUser create user.
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc
	res := &api.UidT{}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("start to create user,arg: %v", u)

	user := &m.User{}
	user.Uid = m.GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return res, handValidataError(c, err)
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to create user")
		return res, e.ErrInternalError
	}
	res.Uid = user.Uid

	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to create user")
	return res, nil
}

// ReadUser read user.
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc
	res := &api.UserT{}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to read user")

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return res, handValidataError(c, err)
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user uid, uid = %v", uid)

	u, err := svc.ReadUser(c, uid.Uid)
	if err != nil {
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", res.Uid).
			Msg("failed to read user")
		return res, e.ErrInternalError
	}

	res.Uid = u.Uid
	res.Name = u.Name
	res.Sex = u.Sex

	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", res.Uid).
		Msg("succ to read user")
	return res, nil
}

// UpdateUser update user.
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("start to update user, arg: %v", u)

	user := &m.User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return empty, handValidataError(c, err)
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("failed to update user")
		return empty, e.ErrInternalError
	}
	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Msg("succ to update user")
	return empty, nil
}

// DeleteUser delete user.
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msg("start to delete user")

	user := &m.User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		return empty, handValidataError(c, err)
	}

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("succ to get user uid, uid = %v", uid)

	err := svc.DeleteUser(c, uid.Uid)
	if err != nil {
		log.Info().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", uid.Uid).
			Msg("failed to delete user")
		return empty, e.ErrInternalError
	}

	log.Info().
		Int64("request_id", rqid.GetIdMust(c)).
		Int64("user_id", uid.Uid).
		Msg("failed to delete user")
	return empty, nil
}
