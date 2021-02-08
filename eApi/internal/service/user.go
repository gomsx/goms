package service

import (
	"context"

	m "github.com/fuwensun/goms/eApi/internal/model"

	log "github.com/sirupsen/logrus"
)

// CreateUser create user.
func (s *service) CreateUser(ctx context.Context, user *m.User) error {
	err := s.dao.CreateUser(ctx, user)
	if err != nil {
		log.Errorf("failed to create user: %v", err)
		return err
	}
	return nil
}

// ReadUser read user.
func (s *service) ReadUser(ctx context.Context, uid int64) (*m.User, error) {
	user, err := s.dao.ReadUser(ctx, uid)
	if err != nil {
		log.Errorf("failed to read user: %v", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser update user.
func (s *service) UpdateUser(ctx context.Context, user *m.User) error {
	err := s.dao.UpdateUser(ctx, user)
	if err != nil {
		log.Errorf("failed to update user: %v", err)
		return err
	}
	return nil
}

// DeleteUser delete user.
func (s *service) DeleteUser(ctx context.Context, uid int64) error {
	err := s.dao.DeleteUser(ctx, uid)
	if err != nil {
		log.Errorf("failed to delete user: %v", err)
		return err
	}
	return nil
}
