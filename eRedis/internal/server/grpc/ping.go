package grpc

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eRedis/api"
	m "github.com/fuwensun/goms/eRedis/internal/model"
	. "github.com/fuwensun/goms/eRedis/internal/pkg/err"
)

// Ping ping server.
func (s *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	svc := s.svc
	//
	var res *api.Reply
	ping := &m.Ping{Type: "grpc"}
	ping, err := svc.HandPing(c, ping)
	if err != nil {
		res = &api.Reply{
			Message: ErrInternalError.Error(),
		}
		return res, err
	}
	//
	res = &api.Reply{
		Message: m.MakePongMsg(req.Message),
		Count:   ping.Count,
	}
	log.Printf("pong msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
