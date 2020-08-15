package grpc

import (
	"context"
	"log"

	"github.com/aivuca/goms/eMysql/api"
	m "github.com/aivuca/goms/eMysql/internal/model"
	e "github.com/aivuca/goms/eMysql/internal/pkg/err"
)

// Ping ping methon.
func (s *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := s.svc
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
	log.Printf("pong msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
