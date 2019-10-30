package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/fuwensun/goms/eMysql/internal/model"
)

func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func (d *dao) UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) {

	fmt.Printf("dao update ping count: %v => %v\n", t, v)

	db := d.db
	//查询数据
	rows, err := db.Query("select * from api_test_ping_count")
	checkErr(err)

	for rows.Next() { //遍历Rows
		var pt model.PingType
		var pc model.PingCount
		err = rows.Scan(&pt, &pc) //获取一行结果
		checkErr(err)
		fmt.Println(pt, pc)
	}
	defer rows.Close() //释放链接

	//更新数据
	stmt, err := db.Prepare("update api_test_ping_count set  count=? where type=?")
	checkErr(err)

	es := fmt.Sprintf("%s", t)
	res, err := stmt.Exec(v, es)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)
}
