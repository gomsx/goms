package service

import (
	"context"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

//http
func (s *Service) UpdateHttpPingCount(c context.Context, pingcount model.PingCount) error {
	err := s.dao.UpdatePingCount(c, model.HTTP, pingcount)
	return err

}

//
func (s *Service) ReadHttpPingCount(c context.Context) (model.PingCount, error) {
	pc, err := s.dao.ReadPingCount(c, model.HTTP)
	return pc, err
}

//grpc
func (s *Service) UpdateGrpcPingCount(c context.Context, pingcount model.PingCount) error {
	err := s.dao.UpdatePingCount(c, model.GRPC, pingcount)
	return err
}

//
func (s *Service) ReadGrpcPingCount(c context.Context) (model.PingCount, error) {
	pc, err := s.dao.ReadPingCount(c, model.GRPC)
	return pc, err
}
