package service

import (
	"context"

	. "github.com/gomsx/goms/eRedis/internal/model"
)

// CreateUser create user.
func (s *service) CreateUser(ctx context.Context, user *User) error {
	err := s.dao.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// ReadUser read user.
func (s *service) ReadUser(ctx context.Context, uid int64) (*User, error) {
	user, err := s.dao.ReadUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser update user.
func (s *service) UpdateUser(ctx context.Context, user *User) error {
	err := s.dao.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser delete user.
func (s *service) DeleteUser(ctx context.Context, uid int64) error {
	err := s.dao.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}
