package dao

import (
	"context"
	"fmt"

	. "github.com/fuwensun/goms/eApi/internal/model"
)

const (
	_updatePing = "UPDATE ping_table SET count=? WHERE type=?"
	_readPing   = "SELECT count FROM ping_table WHERE type=?"
)

func (d *dao) UpdatePing(c context.Context, p *Ping) error {
	db := d.db
	if _, err := db.Exec(_updatePing, p.Count, p.Type); err != nil {
		err = fmt.Errorf("db exec update: %w", err)
		return err
	}
	return nil
}

func (d *dao) ReadPing(c context.Context, t string) (p *Ping, err error) {
	db := d.db
	p = &Ping{}
	rows, err := db.Query(_readPing, p.Type)
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
	}
	return p, nil
}
