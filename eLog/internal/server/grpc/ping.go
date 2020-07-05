package grpc

import (
	"context"

	"github.com/aivuca/goms/eLog/api"
	. "github.com/aivuca/goms/eLog/internal/model"
	"github.com/rs/zerolog/log"
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
	log.Info().Msgf("grpc ping msg: %v, count: %v", msg, p.Count)
	return res, nil
}
