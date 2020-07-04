package service

import (
	"context"

	"github.com/aivuca/goms/eMysql/internal/model"
)

// http
func (s *Service) HandPingHttp(c context.Context) (model.PingCount, error) {
	dao := s.dao
	pc, err := dao.ReadPingCount(c, model.HTTP)
	if err != nil {
		return pc, err
	}
	pc++
	err = dao.UpdatePingCount(c, model.HTTP, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}

// grpc
func (s *Service) HandPingGrpc(c context.Context) (model.PingCount, error) {
	dao := s.dao
	pc, err := dao.ReadPingCount(c, model.GRPC)
	if err != nil {
		return pc, err
	}
	pc++
	err = dao.UpdatePingCount(c, model.GRPC, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}

