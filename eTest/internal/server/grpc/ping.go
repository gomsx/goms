package grpc

import (
	"context"

	"github.com/fuwensun/goms/eTest/api"
	m "github.com/fuwensun/goms/eTest/internal/model"
	e "github.com/fuwensun/goms/eTest/internal/pkg/err"

	"github.com/rs/zerolog/log"
)

// Ping
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
	res = &api.Reply{
		Message: m.MakePongMsg(req.Message),
		Count:   p.Count,
	}
	log.Debug().
		Msgf("ping msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
