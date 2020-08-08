package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"
)

// Ping.
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := srv.svc
	p := &m.Ping{}
	p.Type = "grpc"
	p, err := svc.HandPing(c, p)
	if err != nil {
		res = &api.Reply{
			Message: e.ErrInternalError.Error(),
		}
		return res, err
	}
	msg := "Pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		Count:   p.Count,
	}
	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("ping msg: %v, count: %v", msg, p.Count)
	return res, nil
}
