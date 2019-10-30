package dao

import (
	"context"
	"fmt"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

func (d *dao) UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error {
	db := d.db
	//更新数据
	stmt, err := db.Prepare("update api_test_ping_count set count=? where type=?")
	if err != nil {
		err = fmt.Errorf("failed to prepare error [%w]", err)
		return err
	}
	_, err = stmt.Exec(v, t)
	if err != nil {
		err = fmt.Errorf("failed to exec error [%w]", err)
		return err
	}
	return nil
}

func (d *dao) ReadPingCount(c context.Context, t model.PingType) (pc model.PingCount, err error) {
	db := d.db
	//查询数据
	rows, err := db.Query(fmt.Sprintf("select count from api_test_ping_count where type='%s'", t))
	if err != nil {
		err = fmt.Errorf("failed to query, error [%w]", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&pc) //获取一行结果
		if err != nil {
			err = fmt.Errorf("failed to scan, error [%w]", err)
			return
		}
	}
	defer rows.Close() //释放链接

	return pc, nil
}
