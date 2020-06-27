package grpc

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/fuwensun/goms/eTest/api"
	. "github.com/fuwensun/goms/eTest/internal/model"
)

var empty = &api.Empty{}

// createUser
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc

	var err error
	res := &api.UidT{}

	log.Debug().Msgf("grpc CreateUser() get arg: %v", u)

	if ok := CheckSex(u.Sex); !ok {
		log.Debug().Msgf("grpc sex err: %v", u.Sex)
		return res, ErrSexError
	}
	if ok := CheckName(u.Name); !ok {
		log.Debug().Msgf("grpc name err: %v", u.Name)
		return res, ErrNameError
	}

	user := User{}
	user.Name = u.Name
	user.Sex = u.Sex
	if err = svc.CreateUser(c, &user); err != nil {
		log.Warn().Msgf("grpc create user: %v", err)
		return res, ErrInternalError
	}

	res.Uid = user.Uid
	log.Info().Msgf("grpc create user=%v", user)
	return res, nil
}

// readUser
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc

	var err error
	user := &api.UserT{}

	if ok := CheckUid(uid.Uid); !ok {
		log.Debug().Msgf("grpc uid err: %v", uid.Uid)
		return user, ErrUidError
	}

	u, err := svc.ReadUser(c, uid.Uid)
	if err == ErrNotFound {
		log.Warn().Msgf("grpc read user: %v", err)
		return user, ErrNotFoundData
	} else if err != nil {
		log.Error().Msgf("grpc read user: %v", err)
		return user, ErrInternalError
	}

	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	log.Info().Msgf("grpc read user=%v", *u)

	return user, nil
}

// updateUser
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc
	var err error

	log.Debug().Msgf("grpc CreateUser() get arg: %v", u)

	if ok := CheckUid(u.Uid); !ok {
		log.Debug().Msgf("grpc uid err: %v", u.Uid)
		return empty, ErrUidError
	}
	if ok := CheckSex(u.Sex); !ok {
		log.Debug().Msgf("grpc sex err: %v", u.Sex)
		return empty, ErrSexError
	}
	if ok := CheckName(u.Name); !ok {
		log.Debug().Msgf("grpc name err: %v", u.Name)
		return empty, ErrNameError
	}

	user := User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	err = svc.UpdateUser(c, &user)
	if err == ErrNotFound {
		log.Warn().Msgf("grpc update user: %v", err)
		return empty, ErrNotFoundData
	} else if err != nil {
		log.Error().Msgf("grpc update user: %v", err)
		return empty, ErrInternalError
	}
	log.Info().Msgf("grpc update user=%v", user)
	return empty, nil
}

// deleteUser
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc
	var err error

	if ok := CheckUid(uid.Uid); !ok {
		log.Debug().Msgf("grpc uid err: %v", uid.Uid)
		return empty, ErrUidError
	}

	err = svc.DeleteUser(c, uid.Uid)
	if err == ErrNotFound {
		log.Warn().Msgf("grpc delete user: %v", err)
		return empty, ErrNotFoundData
	} else if err != nil {
		log.Error().Msgf("grpc delete user: %v", err)
		return empty, ErrInternalError
	}
	log.Info().Msgf("grpc delete user uid=%v", uid.Uid)
	return empty, nil
}
