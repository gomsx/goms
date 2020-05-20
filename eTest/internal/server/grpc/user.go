package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/fuwensun/goms/eTest/api"
	. "github.com/fuwensun/goms/eTest/internal/model"
)

var empty = &api.Empty{}

// createUser
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc

	var err error
	res := &api.UidT{}

	if ok := CheckSex(u.Sex); !ok {
		log.Printf("grpc sex err: %v", u.Sex)
		err = fmt.Errorf("sex error!")
		return res, err
	}
	if ok := CheckName(u.Name); !ok {
		log.Printf("grpc name err: %v", u.Name)
		err = fmt.Errorf("name error!")
		return res, err
	}

	user := User{}
	user.Name = u.Name
	user.Sex = u.Sex
	if err = svc.CreateUser(c, &user); err != nil {
		log.Printf("grpc create user: %v", err)
		return res, ErrInternalError
	}
	log.Printf("grpc create user=%v", user)
	res.Uid = user.Uid
	return res, nil
}

// updateUser
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc
	var err error

	if ok := CheckUid(u.Uid); !ok {
		log.Printf("grpc uid err: %v", u.Uid)
		err = fmt.Errorf("uid error!")
		return empty, err
	}
	if ok := CheckSex(u.Sex); !ok {
		log.Printf("grpc sex err: %v", u.Sex)
		err = fmt.Errorf("sex error!")
		return empty, err
	}
	if ok := CheckName(u.Name); !ok {
		log.Printf("grpc name err: %v", u.Name)
		err = fmt.Errorf("name error!")
		return empty, err
	}

	user := User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex
	err = svc.UpdateUser(c, &user)
	if err == ErrNotFound {
		log.Printf("grpc update user: %v", err)
		return empty, ErrNotFoundData
	} else if err != nil {
		log.Printf("grpc update user: %v", err)
		return empty, ErrInternalError
	}
	log.Printf("grpc update user=%v", user)
	return empty, nil
}

// readUser
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc

	var err error
	user := &api.UserT{}

	if ok := CheckUid(uid.Uid); !ok {
		log.Printf("grpc uid err: %v", uid.Uid)
		err = fmt.Errorf("uid error!")
		return user, err
	}

	u, err := svc.ReadUser(c, uid.Uid)
	if err == ErrNotFound {
		log.Printf("grpc read user: %v", err)
		return user, ErrNotFoundData
	} else if err != nil {
		log.Printf("grpc read user: %v", err)
		return user, ErrInternalError
	}

	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex
	log.Printf("grpc read user=%v", u)
	return user, nil
}

// deleteUser
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc
	var err error

	if ok := CheckUid(uid.Uid); !ok {
		log.Printf("grpc uid err: %v", uid.Uid)
		err = fmt.Errorf("uid error!")
		return empty, err
	}

	err = svc.DeleteUser(c, uid.Uid)
	if err == ErrNotFound {
		log.Printf("grpc delete user: %v", err)
		return empty, ErrNotFoundData
	} else if err != nil {
		log.Printf("grpc delete user: %v", err)
		return empty, ErrInternalError
	}
	log.Printf("grpc delete user uid=%v", uid.Uid)
	return empty, nil
}
