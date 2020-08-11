package service

import (
	"context"

	m "github.com/aivuca/goms/eApi/internal/model"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"
)

// CreateUser create user.
func (s *service) CreateUser(c context.Context, user *m.User) error {
	err := s.dao.CreateUser(c, user)
	if err != nil {
		log.Error().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("failed to create user, err = %v", err)
		return err
	}
	return nil
}

// ReadUser read user.
func (s *service) ReadUser(c context.Context, uid int64) (*m.User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if err != nil {
		log.Error().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("failed to read user, err = %v", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser update user.
func (s *service) UpdateUser(c context.Context, user *m.User) error {
	err := s.dao.UpdateUser(c, user)
	if err != nil {
		log.Error().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("failed to update user, err = %v", err)
		return err
	}
	return nil
}

// DeleteUser delete user.
func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if err != nil {
		log.Error().
			Int64("request_id", rqid.GetIdMust(c)).
			Msgf("failed to delete user, err = %v", err)
		return err
	}
	return nil
}
