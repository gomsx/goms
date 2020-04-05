package grpc

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eTest/api"
)

// Ping
func (s *Server) Ping(c context.Context, req *api.Request) (res *api.Reply, err error) {
	svc := s.svc
	pc, err := svc.ReadGrpcPingCount(c)
	if err != nil {
		res = &api.Reply{
			Message: "internal error!",
		}
		return
	}
	pc++
	err = svc.UpdateGrpcPingCount(c, pc)
	if err != nil {
		res = &api.Reply{
			Message: "internal error!",
		}
		return
	}

	msg := "pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		// Count:   pc,
	}
	log.Printf("grpc ping msg: %v count: %v", msg, pc)
	return res, nil
}
