package service

import (
	"context"

	. "github.com/aivuca/goms/eRedis/internal/model"
)

// http
func (s *service) HandPingHttp(c context.Context) (PingCount, error) {
	dao := s.dao
	pc, err := dao.ReadPingCount(c, HTTP)
	if err != nil {
		return pc, err
	}
	pc++
	err = dao.UpdatePingCount(c, HTTP, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}

// grpc
func (s *service) HandPingGrpc(c context.Context) (PingCount, error) {
	dao := s.dao
	pc, err := dao.ReadPingCount(c, GRPC)
	if err != nil {
		return pc, err
	}
	pc++
	err = dao.UpdatePingCount(c, GRPC, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}
