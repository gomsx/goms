package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	. "github.com/aivuca/goms/eApi/internal/model"
)

// Ping
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := srv.svc
	p := &Ping{}
	p.Type = "grpc"
	p, err := svc.HandPing(c, p)
	if err != nil {
		res = &api.Reply{
			Message: ErrInternalError.Error(),
		}
		return res, err
	}
	msg := "Pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		Count:   p.Count,
	}
	log.Debug().Msgf("ping msg: %v, count: %v", msg, p.Count)
	return res, nil
}
