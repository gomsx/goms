package service

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

//
func (s *Service) UpdateHttpPingCount(c context.Context, pingcount model.PingCount) {
	if err := s.dao.UpdatePingCount(c, model.HTTP, pingcount); err != nil {
		log.Fatalf("failed to update http ping count %v", err)
	}
}

//
func (s *Service) ReadHttpPingCount(c context.Context) model.PingCount {
	pc, err := s.dao.ReadPingCount(c, model.HTTP)
	if err != nil {
		log.Fatalf("failed to read http ping count %v", err)
	}
	return pc
}
