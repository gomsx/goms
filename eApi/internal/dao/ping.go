package dao

import (
	"context"
	"fmt"

	. "github.com/fuwensun/goms/eApi/internal/model"
)

const (
	_readPing   = "SELECT count FROM ping_table WHERE type=?"
	_updatePing = "UPDATE ping_table SET count=? WHERE type=?"
)

func (d *dao) ReadPing(c context.Context, t string) (p *Ping, err error) {
	db := d.db
	p = &Ping{}
	rows, err := db.Query(_readPing, t)
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("db query: %w", err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&p.Count)
		if err != nil {
			err = fmt.Errorf("db rows scan: %w", err)
			return
		}
		p.Type = t
		log.Debug().Msgf("db read ping = %v", *p)
		return
	}
	log.Debug().Msgf("db not found ping, type = %v", t)
	err = nil
	return
}

func (d *dao) UpdatePing(c context.Context, p *Ping) error {
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
	log.Debug().Str("type", p.Type).Msg("db update user")
	log.Debug().Int64("rows", num).Msg("db update user")
	return nil
}
