package grpc

import (
	"context"

	"github.com/aivuca/goms/eLog/api"
	m "github.com/aivuca/goms/eLog/internal/model"
	e "github.com/aivuca/goms/pkg/err"
	ms "github.com/aivuca/goms/pkg/misc"

	"github.com/rs/zerolog/log"
)

// Ping ping server.
func (s *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := s.svc
	//
	ping := &m.Ping{Type: "grpc"}
	ping, err := svc.HandPing(c, ping)
	if err != nil {
		res = &api.Reply{
			Message: e.ErrInternalError.Error(),
		}
		return res, err
	}
	//
	res = &api.Reply{
		Message: ms.MakePongMsg(req.Message),
		Count:   ping.Count,
	}
	log.Ctx(c).Debug().
		Msgf("pong msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
