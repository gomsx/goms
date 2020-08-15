package grpc

import (
	"context"

	"github.com/aivuca/goms/eLog/api"
	m "github.com/aivuca/goms/eLog/internal/model"
	e "github.com/aivuca/goms/eLog/internal/pkg/err"

	"github.com/rs/zerolog/log"
)

// Ping ping server.
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := srv.svc
	//
	p := &m.Ping{}
	p.Type = "grpc"

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
