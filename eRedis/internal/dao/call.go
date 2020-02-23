package dao

import (
	"context"
	"fmt"

	"github.com/fuwensun/goms/eRedis/internal/model"
)

func (d *dao) UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error {
	db := d.db
	//更新数据
	stmt, err := db.Prepare("UPDATE api_call_ping_count SET count=? WHERE type=?")
	if err != nil {
		err = fmt.Errorf("prepare db: %w", err)
		return err
	}
	_, err = stmt.Exec(v, t)
	if err != nil {
		err = fmt.Errorf("exec stmt: %w", err)
		return err
	}
	return nil
}

func (d *dao) ReadPingCount(c context.Context, t model.PingType) (pc model.PingCount, err error) {
	db := d.db
	//查询数据
	rows, err := db.Query(fmt.Sprintf("SELECT count FROM api_call_ping_count WHERE type='%s'", t))
	if err != nil {
		err = fmt.Errorf("query db: %w", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&pc) //获取一行结果
		if err != nil {
			err = fmt.Errorf("scan rows: %w", err)
			return
		}
	}
	defer rows.Close() //释放链接

	return pc, nil
}
