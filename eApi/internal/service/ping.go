package service

import (
	"context"

	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/rs/zerolog/log"
)

// HandPing hand ping.
func (s *service) HandPing(c context.Context, p *m.Ping) (*m.Ping, error) {
	dao := s.dao
	p, err := dao.ReadPing(c, p.Type)
	if err != nil {
		log.Ctx(c).Error().
			Msgf("failed to read ping: %v", err)
		return nil, err
	}
	p.Count++
	err = dao.UpdatePing(c, p)
	if err != nil {
		log.Ctx(c).Error().
			Msgf("failed to update ping: %v", err)
		return nil, err
	}
	return p, nil
}
