package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"

	"github.com/rs/zerolog/log"
)

// setPingReplyMate set mate data to ping reply.
func setPingReplyMate(r *api.PingReply, ecode int64, err error) {
	r.Code = ecode
	if err != nil {
		r.Msg = err.Error()
	}
	r.Msg = "ok"
}

// Ping ping server.
func (srv *Server) Ping(c context.Context, in *api.PingReq) (*api.PingReply, error) {
	svc := srv.svc
	res := &api.PingReply{Data: &api.PingMsg{}}
	d := in.Data
	//
	p := &m.Ping{}
	p.Type = "grpc"

	p, err := svc.HandPing(c, p)
	if err != nil {
		setPingReplyMate(res, e.StatusInternalServerError, err)
		return res, err
	}
	//
	res.Data.Message = m.MakePongMsg(d.Message)
	res.Data.Count = p.Count
	setPingReplyMate(res, e.StatusOK, nil)
	log.Ctx(c).Debug().
		Msgf("ping msg: %v, count: %v", res.Data.Message, res.Data.Count)
	return res, nil
}
