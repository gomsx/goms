package dao

import (
	"context"
	"fmt"

	m "github.com/aivuca/goms/eTest/internal/model"

	log "github.com/sirupsen/logrus"
)

const (
	_readPing   = "SELECT type,count FROM ping_table WHERE type=?"
	_updatePing = "UPDATE ping_table SET count=? WHERE type=?"
)

// ReadPing read ping.
func (d *dao) ReadPing(ctx context.Context, t string) (*m.Ping, error) {
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
		log.Debugf("db read ping: %v", *p)
		return p, nil
	}
	log.Debugf("db not found ping, type: %v", t)
	return p, nil //not found data
}

// UpdatePing update ping.
func (d *dao) UpdatePing(ctx context.Context, p *m.Ping) error {
	db := d.db
	result, err := db.Exec(_updatePing, p.Count, p.Type)
	if err != nil {
		err = fmt.Errorf("db exec update: %w", err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	log.Debugf("db update ping: %v", *p)
	return nil
}
