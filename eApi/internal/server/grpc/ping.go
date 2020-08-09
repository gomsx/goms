package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"
)

//
func makeMessage(s string) string {
	return "Pong" + " " + s
}

// Ping.
func (srv *Server) Ping(c context.Context, in *api.PingReq) (*api.PingReply, error) {
	svc := srv.svc
	res := &api.PingReply{Data: &api.PingMsg{}}
	d := in.Data

	p := &m.Ping{}
	p.Type = "grpc"

	p, err := svc.HandPing(c, p)
	if err != nil {
		res = &api.PingReply{
			Code: 500,
			Msg:  e.ErrInternalError.Error(),
		}
		return res, err
	}

	msg := makeMessage(d.Message)
	res.Data = &api.PingMsg{
		Message: msg,
		Count:   p.Count,
	}
	res.Code = 200
	res.Msg = "ok"

	log.Debug().
		Int64("request_id", rqid.GetIdMust(c)).
		Msgf("ping msg: %v, count: %v", msg, p.Count)
	return res, nil
}
