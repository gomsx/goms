package service

import (
	"context"

	"github.com/rs/zerolog/log"
	. "github.com/fuwensun/goms/eLog/internal/model"
	"golang.org/x/exp/errors"
)

func (s *service) CreateUser(c context.Context, user *User) error {
	user.Uid = GetUid()
	err := s.dao.CreateUser(c, user)
	if errors.Is(err, ErrFailedCreateData) {
		log.Warn().Msg("delete user,not found data")
		return err
	} else if err != nil {
		log.Error().Msg("failed to create user")
		return err
	}
	return nil
}

func (s *service) ReadUser(c context.Context, uid int64) (*User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if errors.Is(err, ErrNotFoundData) {
		log.Warn().Msg("delete user,not found data")
		return nil, err
	} else if err != nil {
		log.Error().Msg("failed to read user")
		return nil, err
	}
	return user, nil
}

func (s *service) UpdateUser(c context.Context, user *User) error {
	err := s.dao.UpdateUser(c, user)
	if errors.Is(err, ErrNotFoundData) {
		log.Warn().Msg("delete user,not found data")
		return err
	} else if err != nil {
		log.Error().Msg("failed to update user")
		return err
	}
	return nil
}

func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if errors.Is(err, ErrNotFoundData) {
		log.Warn().Msg("delete user,not found data")
		return err
	} else if err != nil {
		log.Error().Msg("failed to delete user")
		return err
	}
	return nil
}
