package dao

import (
	"context"
	"fmt"
)

func (d *dao) UpdateUserName(c context.Context, uid int64, name string) error {
	db := d.db
	//更新数据
	stmt, err := db.Prepare("UPDATE user_info SET name=? WHERE uid=?")
	if err != nil {
		err = fmt.Errorf("prepare db: %w", err)
		return err
	}
	_, err = stmt.Exec(name, uid)
	if err != nil {
		err = fmt.Errorf("exec stmt: %w", err)
		return err
	}
	return nil

	// fmt.Printf("update user name")
	// return nil
}

func (d *dao) ReadUserName(c context.Context, uid int64) (name string, err error) {
	db := d.db
	//查询数据
	rows, err := db.Query(fmt.Sprintf("SELECT name FROM user_info WHERE uid ='%v'", uid))
	if err != nil {
		err = fmt.Errorf("query db: %w", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&name) //获取一行结果
		if err != nil {
			err = fmt.Errorf("scan rows: %w", err)
			return
		}
	}
	defer rows.Close() //释放链接
	return name, nil
	// fmt.Printf("read user name")
	// return "haha", nil

}
