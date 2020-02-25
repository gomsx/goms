package service

import (
	"context"
	"log"
)

//
func (s *Service) UpdateUserName(c context.Context, uid int64, name string) {
	if err := s.dao.UpdateUserName(c, uid, name); err != nil {
		log.Fatalf("failed to update user name: %v", err)
	}
}

//
func (s *Service) ReadHttpUserName(c context.Context, uid int64) string {
	name, err := s.dao.ReadUserName(c, uid)
	if err != nil {
		log.Fatalf("failed to read user name: %v", err)
	}
	return name
}
