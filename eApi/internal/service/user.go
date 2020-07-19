package service

import (
	"context"

	. "github.com/fuwensun/goms/eApi/internal/model"
)

func (s *service) CreateUser(c context.Context, user *User) error {
	err := s.dao.CreateUser(c, user)
	if err != nil {
		log.Error().Msgf("failed to create user, err = %v", err)
		return err
	}
	return nil
}

func (s *service) ReadUser(c context.Context, uid int64) (*User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if err != nil {
		log.Error().Msgf("failed to read user, err = %v", err)
		return nil, err
	}
	return user, nil
}

func (s *service) UpdateUser(c context.Context, user *User) error {
	err := s.dao.UpdateUser(c, user)
	if err != nil {
		log.Error().Msgf("failed to update user, err = %v", err)
		return err
	}
	return nil
}

func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if err != nil {
		log.Error().Msgf("failed to delete user, err = %v", err)
		return err
	}
	return nil
}
