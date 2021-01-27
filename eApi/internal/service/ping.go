package service

import (
	"context"

	m "github.com/aivuca/goms/eApi/internal/model"

	log "github.com/sirupsen/logrus"
)

// HandPing hand ping.
func (s *service) HandPing(c context.Context, p *m.Ping) (*m.Ping, error) {
	dao := s.dao
	p, err := dao.ReadPing(c, p.Type)
	if err != nil {
		log.Errorf("failed to read ping: %v", err)
		return nil, err
	}
	p.Count++
	err = dao.UpdatePing(c, p)
	if err != nil {
		log.Errorf("failed to update ping: %v", err)
		return nil, err
	}
	return p, nil
}
