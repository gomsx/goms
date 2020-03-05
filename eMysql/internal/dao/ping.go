package dao

import (
	"context"
	"fmt"

	"github.com/fuwensun/goms/eRedis/internal/model"
)

func (d *dao) UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error {
	db := d.db
	if _, err := db.Exec("UPDATE ping_table SET count=? WHERE type=?", v, t); err != nil {
		err = fmt.Errorf("mysql exec update: %w", err)
		return err
	}
	return nil
}

func (d *dao) ReadPingCount(c context.Context, t model.PingType) (pc model.PingCount, err error) {
	db := d.db
	rows, err := db.Query("SELECT count FROM ping_table WHERE type=?", t)
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("mysql query: %w", err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&pc)
		if err != nil {
			err = fmt.Errorf("mysql rows scan: %w", err)
			return
		}
	}
	return pc, nil
}
