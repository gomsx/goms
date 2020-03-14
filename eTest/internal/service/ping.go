package service

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eTest/internal/model"
)

//http
func (s *service) UpdateHttpPingCount(c context.Context, pingcount model.PingCount) {
	if err := s.dao.UpdatePingCount(c, model.HTTP, pingcount); err != nil {
		log.Fatalf("failed to update count of http ping: %v", err)
	}
}

//
func (s *service) ReadHttpPingCount(c context.Context) model.PingCount {
	pc, err := s.dao.ReadPingCount(c, model.HTTP)
	if err != nil {
		log.Fatalf("failed to read count of http ping: %v", err)
	}
	return pc
}

//grpc
func (s *service) UpdateGrpcPingCount(c context.Context, pingcount model.PingCount) {
	if err := s.dao.UpdatePingCount(c, model.GRPC, pingcount); err != nil {
		log.Fatalf("failed to update count of grpc ping: %v", err)
	}
}

//
func (s *service) ReadGrpcPingCount(c context.Context) model.PingCount {
	pc, err := s.dao.ReadPingCount(c, model.GRPC)
	if err != nil {
		log.Fatalf("failed to read count of grpc ping: %v", err)
	}
	return pc
}
