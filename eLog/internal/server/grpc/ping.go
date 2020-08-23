package grpc

import (
	"context"

	"github.com/aivuca/goms/eLog/api"
	m "github.com/aivuca/goms/eLog/internal/model"
	e "github.com/aivuca/goms/eLog/internal/pkg/err"

	"github.com/rs/zerolog/log"
)

// Ping ping server.
func (s *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := s.svc
	//
	p := &m.Ping{Type: "grpc"}

	p, err := svc.HandPing(c, p)
	if err != nil {
		res = &api.Reply{
			Message: e.ErrInternalError.Error(),
		}
		return res, err
	}
	//
	res = &api.Reply{
		Message: m.MakePongMsg(req.Message),
		Count:   p.Count,
	}
	log.Ctx(c).Debug().
		Msgf("pong msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
