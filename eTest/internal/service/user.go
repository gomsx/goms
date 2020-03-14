package service

import (
	"context"
	"log"
	"math/rand"

	"github.com/fuwensun/goms/eTest/internal/model"
	"golang.org/x/exp/errors"
)

func (s *service) CreateUser(c context.Context, user *model.User) error {
	user.Uid = rand.Int63n(0x0FFF_FFFF_FFFF_FFFF) //0x0FFF_FFFF
	err := s.dao.CreateUser(c, user)
	if errors.Is(err, model.ErrFailedCreateData) {
		return err
	} else if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	return nil
}

func (s *service) UpdateUser(c context.Context, user *model.User) error {
	err := s.dao.UpdateUser(c, user)
	if errors.Is(err, model.ErrNotFound) {
		return err
	} else if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}
	return nil
}

func (s *service) ReadUser(c context.Context, uid int64) (model.User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if errors.Is(err, model.ErrNotFound) {
		return user, err
	} else if err != nil {
		log.Fatalf("failed to read user: %v", err)
	}
	return user, nil
}
func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if errors.Is(err, model.ErrNotFound) {
		return err
	} else if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}
	return nil
}
