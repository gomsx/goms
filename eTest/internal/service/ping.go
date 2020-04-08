package service

import (
	"context"

	. "github.com/fuwensun/goms/eTest/internal/model"
)

//http
func (s *service) UpdateHttpPingCount(c context.Context, pingcount PingCount) error {
	err := s.dao.UpdatePingCount(c, HTTP, pingcount)
	return err

}

//
func (s *service) ReadHttpPingCount(c context.Context) (PingCount, error) {
	pc, err := s.dao.ReadPingCount(c, HTTP)
	return pc, err
}

//grpc
func (s *service) UpdateGrpcPingCount(c context.Context, pingcount PingCount) error {
	err := s.dao.UpdatePingCount(c, GRPC, pingcount)
	return err
}

//
func (s *service) ReadGrpcPingCount(c context.Context) (PingCount, error) {
	pc, err := s.dao.ReadPingCount(c, GRPC)
	return pc, err
}
