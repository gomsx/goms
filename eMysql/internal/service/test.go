package service

import (
	"context"
	"fmt"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

//
func (s *Service) HandHttpPingCount(c context.Context, pingcount model.PingCount) {
	fmt.Printf("service http ping count: %v\n", pingcount)
	s.dao.UpdatePingCount(c, model.HTTP, pingcount)
}
