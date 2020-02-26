package service

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"golang.org/x/exp/errors"
)

//
func (s *Service) UpdateUserName(c context.Context, uid int64, name string) {
	if err := s.dao.UpdateUserName(c, uid, name); err != nil {
		log.Fatalf("failed to update user name: %v", err)
	}
}

//
func (s *Service) ReadUserName(c context.Context, uid int64) (string, error) {
	name, err := s.dao.ReadUserName(c, uid)
	if errors.Is(err, model.ErrNotFound) {
		return "", err
	} else if err != nil {
		log.Fatalf("failed to read user name: %v", err)
	}
	return name, nil
}
