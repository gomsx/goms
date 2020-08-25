package grpc

import (
	"context"
	"log"

	"github.com/aivuca/goms/eRedis/api"
	. "github.com/aivuca/goms/eRedis/internal/model"
	. "github.com/aivuca/goms/eRedis/internal/pkg/err"
)

// Ping
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	svc := srv.svc
	//
	var res *api.Reply
	p := &Ping{}
	p.Type = "grpc"
	p, err := svc.HandPing(c, p)
	if err != nil {
		res = &api.Reply{
			Message: ErrInternalError.Error(),
		}
		return res, err
	}
	//
	res = &api.Reply{
		Message: MakePongMsg(req.Message),
		Count:   p.Count,
	}
	log.Printf("ping msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
