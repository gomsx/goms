package service

import (
	"context"

	. "github.com/aivuca/goms/eRedis/internal/model"
)

func (s *service) CreateUser(c context.Context, user *User) error {
	err := s.dao.CreateUser(c, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadUser(c context.Context, uid int64) (*User, error) {
	user, err := s.dao.ReadUser(c, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) UpdateUser(c context.Context, user *User) error {
	err := s.dao.UpdateUser(c, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(c context.Context, uid int64) error {
	err := s.dao.DeleteUser(c, uid)
	if err != nil {
		return err
	}
	return nil
}
