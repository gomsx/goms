package grpc

import (
	"context"

	"github.com/fuwensun/goms/eTest/api"
	m "github.com/fuwensun/goms/eTest/internal/model"
	e "github.com/fuwensun/goms/pkg/err"

	log "github.com/sirupsen/logrus"
)

// Ping ping server.
func (s *Server) Ping(ctx context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := s.svc
	//
	ping := &m.Ping{Type: "grpc"}
	ping, err := svc.HandPing(ctx, ping)
	if err != nil {
		res = &api.Reply{
			Message: e.ErrInternalError.Error(),
		}
		log.Infof("failed to hand ping, error: %v", err)
		return res, err
	}
	//
	res = &api.Reply{
		Message: m.MakePongMsg(req.Message),
		Count:   ping.Count,
	}
	log.Debugf("pong msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
