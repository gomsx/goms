package service

import (
	"context"

	m "github.com/aivuca/goms/eMysql/internal/model"
)

// HandPing hand ping.
func (s *Service) HandPing(ctx context.Context, p *m.Ping) (*m.Ping, error) {
	dao := s.dao
	p, err := dao.ReadPing(ctx, p.Type)
	if err != nil {
		return nil, err
	}
	p.Count++
	err = dao.UpdatePing(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
