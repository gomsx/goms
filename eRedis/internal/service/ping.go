package service

import (
	"context"

	. "github.com/fuwensun/goms/eRedis/internal/model"
)

func (s *service) HandPing(c context.Context, p *Ping) (*Ping, error) {
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
