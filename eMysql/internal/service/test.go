package service

import (
	"context"
	"fmt"
	"log"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

//
func (s *Service) UpdateHttpPingCount(c context.Context, pingcount model.PingCount) {
	fmt.Printf("service http ping count: %v\n", pingcount)
	s.dao.UpdatePingCount(c, model.HTTP, pingcount)
}

func (s *Service) ReadHttpPingCount(c context.Context) model.PingCount {
	// fmt.Printf("service http ping count: %v\n", pingcount)
	pc, err := s.dao.ReadPingCount(c, model.HTTP)
	if err != nil {
		log.Fatalf("failed to read http ping count %v", err)
	}
	return pc
}
