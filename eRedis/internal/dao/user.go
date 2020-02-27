package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/gomodule/redigo/redis"
)

func (d *dao) SetUserCache(c context.Context, user *model.User) error {
	rd := d.redis
	key := model.GetRedisKey(user.Uid)
	if _, err := rd.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("redis Do HMSET err: %w", err)
		return err
	}
	log.Printf("redis set key = %v value = %v\n", key, user)
	return nil
}

func (d *dao) GetUserCache(c context.Context, uid int64) (model.User, error) {
	rd := d.redis
	user := model.User{}
	key := model.GetRedisKey(uid)
	exists, err := redis.Bool(rd.Do("EXISTS", key))
	// fmt.Printf("redis exists %v err = %v\n", exists, err)
	if err != nil {
		err = fmt.Errorf("redis Do EXISTS err: %w", err)
		return user, err
	}
	if exists == false {
		err = fmt.Errorf("redis key [%v] is no exists err!", key)
		return user, err
	}

	value, err := redis.Values(rd.Do("HGETALL", key))
	// fmt.Printf("redis HGETALL %v err = %v\n", value, err)
	if err != nil {
		err = fmt.Errorf("redis Do HGETALL err: %w", err)
		return user, err
	}
	//???
	// if value == nil {
	// 	err = fmt.Errorf("redis Do HGETALL nil value err!")
	// 	return user, err
	// }
	err = redis.ScanStruct(value, &user)
	if err != nil {
		err = fmt.Errorf("redis ScanStruct err: %w", err)
		return user, err
	}
	log.Printf("redis get key=%v val=%v\n", key, user)
	return user, nil
}

func (d *dao) DelUserCache(c context.Context, uid int64) error {
	rd := d.redis
	key := model.GetRedisKey(uid)
	exists, err := redis.Bool(rd.Do("EXISTS", key))
	// fmt.Printf("redis exists %v err = %v\n", exists, err)
	if err != nil {
		err = fmt.Errorf("redis Do EXISTS err: %w", err)
		return err
	}
	if exists == false {
		err = fmt.Errorf("redis key: %w", model.ErrNotFound)
		return err
	}

	if _, err := rd.Do("DEL", key); err != nil {
		err = fmt.Errorf("redis Do DEL err: %w", err)
		return err
	}
	log.Printf("redis delete key = %v\n", key)
	return nil
}

func (d *dao) CreateUserDB(c context.Context, user *model.User) error {
	db := d.db
	result, err := db.Exec("insert into user_table  values(?,?,?)", user.Uid, user.Name, user.Sex)
	if err != nil {
		err = fmt.Errorf("exec insert db: %w", err)
		return err
	}
	_, err = result.LastInsertId() //???
	log.Printf("mysql insert %v\n", user)
	return nil
}

func (d *dao) UpdateUserDB(c context.Context, user *model.User) error {
	db := d.db
	result, err := db.Exec(fmt.Sprintf("UPDATE user_table set name = '%v' ,sex = '%v' where uid = '%v'", user.Name, user.Sex, user.Uid))
	if err != nil {
		err = fmt.Errorf("exec update: %w", err)
		return err
	}
	_, err = result.RowsAffected()
	log.Printf("mysql update %v\n", user)
	return nil
}

func (d *dao) ReadUserDB(c context.Context, uid int64) (user model.User, err error) {
	db := d.db
	rows, err := db.Query(fmt.Sprintf("SELECT uid,name,sex FROM user_table WHERE uid ='%v'", uid))
	defer rows.Close() //释放链接
	if err != nil {
		err = fmt.Errorf("query db: %w", err)
		return
	}
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Name, &user.Sex) //获取一行结果
		if err != nil {
			err = fmt.Errorf("scan rows: %w", err)
			return
		}
		log.Printf("mysql read %v\n", user)
		return user, nil
	}

	log.Printf("mysql read not exists uid=%v\n", uid)
	return user, fmt.Errorf("mysql data not found:%w", model.ErrNotFound)
}

//
func (d *dao) CreateUser(c context.Context, user *model.User) error {
	err := d.CreateUserDB(c, user)
	if err != nil {
		return err
	}
	return nil
}

//
func (d *dao) UpdateUser(c context.Context, user *model.User) error {
	err := d.UpdateUserDB(c, user)
	if err != nil {
		return err
	}
	d.DelUserCache(c, user.Uid)
	return nil
}
func (d *dao) ReadUser(c context.Context, uid int64) (model.User, error) {
	user, err := d.GetUserCache(c, uid)
	if err == nil {
		return user, nil
	}
	user, err = d.ReadUserDB(c, uid)
	if err != nil {
		return user, err
	}
	d.SetUserCache(c, &user)
	return user, nil
}

//
func (d *dao) ReadUserName(c context.Context, uid int64) (name string, err error) {
	user, err := d.ReadUser(c, uid)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func (d *dao) UpdateUserName(c context.Context, uid int64, name string) error {
	db := d.db
	//更新数据
	stmt, err := db.Prepare("UPDATE user_table SET name=? WHERE uid=?")
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
}
