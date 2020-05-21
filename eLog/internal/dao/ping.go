package dao

import (
	"context"
	"fmt"

	."github.com/fuwensun/goms/eLog/internal/model"
)

const (
	_updatePingCount = "UPDATE ping_table SET count=? WHERE type=?"
	_readPingCount   = "SELECT count FROM ping_table WHERE type=?"
)

func (d *dao) UpdatePingCount(c context.Context, t  PingType, v  PingCount) error {
	db := d.db
	if _, err := db.Exec(_updatePingCount, v, t); err != nil {
		err = fmt.Errorf("db exec update: %w", err)
		return err
	}
	return nil
}

func (d *dao) ReadPingCount(c context.Context, t  PingType) (pc  PingCount, err error) {
	db := d.db
	rows, err := db.Query(_readPingCount, t)
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("db query: %w", err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&pc)
		if err != nil {
			err = fmt.Errorf("db rows scan: %w", err)
			return
		}
	}
	return pc, nil
}
