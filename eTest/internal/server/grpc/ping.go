package grpc

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eTest/api"
)

// Ping
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	svc := srv.svc
	pc, err := svc.HandPingGrpc(c)
	if err != nil {
		res = &api.Reply{
			Message: "internal error!",
		}
		return res, err
	}
	msg := "pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		Count:   int64(pc),
	}
	log.Printf("grpc ping msg: %v, count: %v", msg, pc)
	return res, nil
}
