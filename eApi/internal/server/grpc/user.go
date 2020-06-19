package grpc

import (
	"context"

	"github.com/fuwensun/goms/eApi/api"
	. "github.com/fuwensun/goms/eApi/internal/model"
)

var empty = &api.Empty{}

// createUser
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc
	res := &api.UidT{}

	log.Debug().Msgf("start to create user,arg: %v", u)

	if ok := CheckSex(u.Sex); !ok {
		log.Debug().Msgf("sex err, sex = %v", u.Sex)
		return res, ErrSexError
	}
	if ok := CheckName(u.Name); !ok {
		log.Debug().Msgf("name err, name = %v", u.Name)
		return res, ErrNameError
	}

	user := &User{}
	user.Name = u.Name
	user.Sex = u.Sex

	log.Debug().Msgf("succ to get user data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		log.Info().Int64("uid", user.Uid).Msg("failed to create user")
		return res, ErrInternalError
	}
	res.Uid = user.Uid

	log.Info().Int64("uid", user.Uid).Msg("succ to create user")
	return res, nil
}

// readUser
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc
	res := &api.UserT{}

	log.Debug().Msg("start to read user")

	if ok := CheckUid(uid.Uid); !ok {
		log.Debug().Msgf("uid err, uid = %v", uid.Uid)
		return res, ErrUidError
	}

	log.Debug().Msgf("succ to get user uid, uid = %v", uid)

	u, err := svc.ReadUser(c, uid.Uid)
	if err != nil {
		log.Info().Int64("uid", res.Uid).Msg("failed to read user")
		return res, ErrInternalError
	}

	res.Uid = u.Uid
	res.Name = u.Name
	res.Sex = u.Sex

	log.Info().Int64("uid", res.Uid).Msg("succ to read user")
	return res, nil
}

// updateUser
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc

	log.Debug().Msgf("start to update user, arg: %v", u)

	if ok := CheckUid(u.Uid); !ok {
		log.Debug().Msgf("uid err, err = %v", u.Uid)
		return empty, ErrUidError
	}
	if ok := CheckSex(u.Sex); !ok {
		log.Debug().Msgf("sex err, err = %v", u.Sex)
		return empty, ErrSexError
	}
	if ok := CheckName(u.Name); !ok {
		log.Debug().Msgf("name err, err = %v", u.Name)
		return empty, ErrNameError
	}

	user := &User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	log.Debug().Msgf("succ to get user data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		log.Info().Int64("uid", user.Uid).Msg("failed to update user")
		return empty, ErrInternalError
	}
	log.Info().Int64("uid", user.Uid).Msg("succ to update user")
	return empty, nil
}

// deleteUser
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc

	log.Debug().Msg("start to delete user")

	if ok := CheckUid(uid.Uid); !ok {
		log.Debug().Msgf("uid err, err = %v", uid.Uid)
		return empty, ErrUidError
	}

	log.Debug().Msgf("succ to get user uid, uid = %v", uid)

	err := svc.DeleteUser(c, uid.Uid)
	if err != nil {
		log.Info().Int64("uid", uid.Uid).Msg("failed to delete user")
		return empty, ErrInternalError
	}

	log.Info().Int64("uid", uid.Uid).Msg("failed to delete user")
	return empty, nil
}
