package grpc

import (
	"context"

	"github.com/fuwensun/goms/eLog/api"
	. "github.com/fuwensun/goms/eLog/internal/model"
	"github.com/rs/zerolog/log"
)

var empty = &api.Empty{}

// createUser
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc

	var err error
	res := &api.UidT{}

	if ok := CheckSex(u.Sex); !ok {
		log.Info().Msgf("grpc sex err: %v", u.Sex)
		return res, ErrSexError
	}
	if ok := CheckName(u.Name); !ok {
		log.Info().Msgf("grpc name err: %v", u.Name)
		return res, ErrNameError
	}

	user := User{}
	user.Name = u.Name
	user.Sex = u.Sex
	if err = svc.CreateUser(c, &user); err != nil {
		log.Info().Msgf("grpc create user: %v", err)
		return res, ErrInternalError
	}
	log.Info().Msgf("grpc create user=%v", user)
	res.Uid = user.Uid
	return res, nil
}

// updateUser
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc
	var err error

	if ok := CheckUid(u.Uid); !ok {
		log.Info().Msgf("grpc uid err: %v", u.Uid)
		return empty, ErrUidError
	}
	if ok := CheckSex(u.Sex); !ok {
		log.Info().Msgf("grpc sex err: %v", u.Sex)
		return empty, ErrSexError
	}
	if ok := CheckName(u.Name); !ok {
		log.Info().Msgf("grpc name err: %v", u.Name)
		return empty, ErrNameError
	}

	user := User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex
	err = svc.UpdateUser(c, &user)
	if err == ErrNotFound {
		log.Info().Msgf("grpc update user: %v", err)
		return empty, ErrNotFoundData
	} else if err != nil {
		log.Info().Msgf("grpc update user: %v", err)
		return empty, ErrInternalError
	}
	log.Info().Msgf("grpc update user=%v", user)
	return empty, nil
}

// readUser
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc

	var err error
	user := &api.UserT{}

	if ok := CheckUid(uid.Uid); !ok {
		log.Info().Msgf("grpc uid err: %v", uid.Uid)
		return user, ErrUidError
	}

	u, err := svc.ReadUser(c, uid.Uid)
	if err == ErrNotFound {
		log.Info().Msgf("grpc read user: %v", err)
		return user, ErrNotFoundData
	} else if err != nil {
		log.Info().Msgf("grpc read user: %v", err)
		return user, ErrInternalError
	}

	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex
	log.Info().Msgf("grpc read user=%v", u)
	return user, nil
}

// deleteUser
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc
	var err error

	if ok := CheckUid(uid.Uid); !ok {
		log.Info().Msgf("grpc uid err: %v", uid.Uid)
		return empty, ErrUidError
	}

	err = svc.DeleteUser(c, uid.Uid)
	if err == ErrNotFound {
		log.Info().Msgf("grpc delete user: %v", err)
		return empty, ErrNotFoundData
	} else if err != nil {
		log.Info().Msgf("grpc delete user: %v", err)
		return empty, ErrInternalError
	}
	log.Info().Msgf("grpc delete user uid=%v", uid.Uid)
	return empty, nil
}
