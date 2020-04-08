package grpc

import (
	"context"
	"log"

	"github.com/fuwensun/goms/eTest/api"
	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/eTest/internal/service"
)

// Ping
func (srv *Server) Ping(c context.Context, req *api.Request) (*api.Reply, error) {
	var res *api.Reply
	pc, err := handping(c, srv.svc)
	if err != nil {
		res = &api.Reply{
			Message: "internal error!",
		}
		return res, err
	}
	msg := "pong" + " " + req.Message
	res = &api.Reply{
		Message: msg,
		// Count:   pc,
	}
	log.Printf("grpc ping msg: %v count: %v", msg, pc)
	return res, nil
}

// hangping
func handping(c context.Context, svc service.Svc) (model.PingCount, error) {
	pc, err := svc.ReadGrpcPingCount(c)
	if err != nil {
		return pc, err
	}
	pc++
	err = svc.UpdateGrpcPingCount(c, pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}
