package service

import (
	"context"

	m "github.com/fuwensun/goms/eRedis/internal/model"
)

// HandPing hand ping.
func (s *service) HandPing(c context.Context, p *m.Ping) (*m.Ping, error) {
	dao := s.dao
	p, err := dao.ReadPing(c, p.Type)
	if err != nil {
		return nil, err
	}
	p.Count++
	err = dao.UpdatePing(c, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
