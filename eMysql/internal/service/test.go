package service

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

//http
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

//grpc
func (s *Service) UpdateGrpcPingCount(c context.Context, pingcount model.PingCount) {
	if err := s.dao.UpdatePingCount(c, model.GRPC, pingcount); err != nil {
		log.Fatalf("failed to update grpc ping count %v", err)
	}
}

//
func (s *Service) ReadGrpcPingCount(c context.Context) model.PingCount {
	pc, err := s.dao.ReadPingCount(c, model.GRPC)
	if err != nil {
		log.Fatalf("failed to read grpc ping count %v", err)
	}
	return pc
}
