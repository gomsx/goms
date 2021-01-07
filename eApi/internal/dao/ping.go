package dao

import (
	"context"
	"fmt"

	m "github.com/aivuca/goms/eApi/internal/model"

	"github.com/rs/zerolog/log"
)

const (
	_readPing   = "SELECT type,count FROM ping_table WHERE type=?"
	_updatePing = "UPDATE ping_table SET count=? WHERE type=?"
)

// ReadPing read ping.
func (d *dao) ReadPing(c context.Context, t string) (*m.Ping, error) {
	db := d.db
	p := &m.Ping{}
	rows, err := db.Query(_readPing, t)
	if err != nil {
		err = fmt.Errorf("db query: %w", err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&p.Type, &p.Count)
		if err != nil {
			err = fmt.Errorf("db rows scan: %w", err)
			return nil, err
		}
		log.Ctx(c).Debug().
			Msgf("db read ping: %v", *p)
		return p, nil
	}
	log.Ctx(c).Debug().
		Msgf("db not found ping, type: %v", t)
	return p, nil //not found data
}

// UpdatePing update ping.
func (d *dao) UpdatePing(c context.Context, p *m.Ping) error {
	db := d.db
	result, err := db.Exec(_updatePing, p.Count, p.Type)
	if err != nil {
		err = fmt.Errorf("db exec update: %w", err)
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	log.Ctx(c).Debug().
		Int64("rows", num).
		Msgf("db update ping: %v", *p)
	return nil
}
