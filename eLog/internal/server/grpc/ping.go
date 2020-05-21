package grpc

import (
	"context"

	"github.com/fuwensun/goms/eLog/api"
	. "github.com/fuwensun/goms/eLog/internal/model"

	"github.com/rs/zerolog/log"
)

// Ping
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := srv.svc
	pc, err := svc.HandPingGrpc(c)
	if err != nil {
		res = &api.Reply{
			Message: ErrInternalError.Error(),
		}
		return res, err
	}
	msg := "Pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		Count:   int64(pc),
	}
	log.Info().Msgf("grpc ping msg: %v, count: %v", msg, pc)
	return res, nil
}
