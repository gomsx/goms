package grpc

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eTest/api"
	"github.com/fuwensun/goms/eTest/internal/model"
)

var pingcount model.PingCount

// example for grpc request handler.
func (s *Server) Ping(ctx context.Context, req *api.Request) (res *api.Reply, err error) {
	svc := s.svc

	pingcount++
	err = svc.UpdateGrpcPingCount(ctx, pingcount)
	if err != nil {
		res = &api.Reply{
			Message: "update grpc ping count error!",
		}
		return
	}
	pc, err := svc.ReadGrpcPingCount(ctx)
	if err != nil {
		res = &api.Reply{
			Message: "read grpc ping count error!",
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
