package service

import (
	"context"
	"log"

	. "github.com/fuwensun/goms/eRedis/internal/model"
	"golang.org/x/exp/errors"
)

func (s *service) CreateUser(c context.Context, user *User) error {
	user.Uid = GetUid()
	err := s.dao.CreateUser(c, user)
	if errors.Is(err, ErrFailedCreateData) {
		return err
	} else if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	return nil
}

func (s *service) UpdateUser(c context.Context, user *User) error {
	err := s.dao.UpdateUser(c, user)
	if errors.Is(err, ErrNotFoundData) {
		return err
	} else if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}
	return nil
}

func (s *service) ReadUser(c context.Context, uid int64) (User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if errors.Is(err, ErrNotFoundData) {
		return user, err
	} else if err != nil {
		log.Fatalf("failed to read user: %v", err)
	}
	return user, nil
}
func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if errors.Is(err, ErrNotFoundData) {
		return err
	} else if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}
	return nil
}
