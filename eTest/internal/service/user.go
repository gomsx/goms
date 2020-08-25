package service

import (
	"context"

	m "github.com/fuwensun/goms/eTest/internal/model"

	"github.com/rs/zerolog/log"
)

// CreateUser create user.
func (s *service) CreateUser(c context.Context, user *m.User) error {
	err := s.dao.CreateUser(c, user)
	if err != nil {
		log.Ctx(c).Error().
			Msgf("failed to create user: %v", err)
		return err
	}
	return nil
}

//ReadUser read user.
func (s *service) ReadUser(c context.Context, uid int64) (*m.User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if err != nil {
		log.Ctx(c).Error().
			Msgf("failed to read user: %v", err)
		return nil, err
	}
	return user, nil
}

//UpdateUser update user.
func (s *service) UpdateUser(c context.Context, user *m.User) error {
	err := s.dao.UpdateUser(c, user)
	if err != nil {
		log.Ctx(c).Error().
			Msgf("failed to update user: %v", err)
		return err
	}
	return nil
}

//
func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if err != nil {
		log.Ctx(c).Error().
			Msgf("failed to delete user: %v", err)
		return err
	}
	return nil
}
