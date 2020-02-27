package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/gomodule/redigo/redis"
)

func (d *dao) existUserCache(c context.Context, uid int64) (bool, error) {
	rd := d.redis
	key := model.GetRedisKey(uid)
	exist, err := redis.Bool(rd.Do("EXISTS", key))
	if err != nil {
		err = fmt.Errorf("redis Do EXISTS err: %w", err)
		return exist, err
	}
	log.Printf("redis exist key=%v\n", exist)
	return exist, nil
}

func (d *dao) setUserCache(c context.Context, user *model.User) error {
	rd := d.redis
	key := model.GetRedisKey(user.Uid)
	if _, err := rd.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("redis Do HMSET err: %w", err)
		return err
	}
	log.Printf("redis set key = %v, value = %v\n", key, user)
	return nil
}

func (d *dao) getUserCache(c context.Context, uid int64) (model.User, error) {
	rd := d.redis
	user := model.User{}
	key := model.GetRedisKey(uid)
	value, err := redis.Values(rd.Do("HGETALL", key))
	if err != nil {
		err = fmt.Errorf("redis Do HGETALL err: %w", err)
		return user, err
	}
	err = redis.ScanStruct(value, &user)
	if err != nil {
		err = fmt.Errorf("redis ScanStruct err: %w", err)
		return user, err
	}
	log.Printf("redis get key=%v, val=%v\n", key, user)
	return user, nil
}

func (d *dao) delUserCache(c context.Context, uid int64) error {
	rd := d.redis
	key := model.GetRedisKey(uid)
	if _, err := rd.Do("DEL", key); err != nil {
		err = fmt.Errorf("redis Do DEL err: %w", err)
		return err
	}
	log.Printf("redis delete key = %v\n", key)
	return nil
}

func (d *dao) createUserDB(c context.Context, user *model.User) error {
	db := d.db
	result, err := db.Exec("insert into user_table  values(?,?,?)", user.Uid, user.Name, user.Sex)
	if err != nil {
		err = fmt.Errorf("exec insert db: %w", err)
		return err
	}
	_, err = result.LastInsertId() //???
	log.Printf("mysql insert user = %v\n", user)
	return nil
}

func (d *dao) updateUserDB(c context.Context, user *model.User) error {
	db := d.db
	result, err := db.Exec(fmt.Sprintf("UPDATE user_table set name = '%v' ,sex = '%v' where uid = '%v'", user.Name, user.Sex, user.Uid))
	if err != nil {
		err = fmt.Errorf("exec update: %w", err)
		return err
	}
	num, err := result.RowsAffected()
	//???
	if num == 0 {
		return model.ErrNotFoundData
	}
	log.Printf("mysql update user = %v, affected = %v\n", user, num)
	return nil
}

func (d *dao) readUserDB(c context.Context, uid int64) (model.User, error) {
	db := d.db
	user := model.User{}
	rows, err := db.Query(fmt.Sprintf("SELECT uid,name,sex FROM user_table WHERE uid ='%v'", uid))
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("query db: %w", err)
		return user, err
	}
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Name, &user.Sex)
		if err != nil {
			err = fmt.Errorf("scan rows: %w", err)
			return user, err
		}
		log.Printf("mysql read user %v\n", user)
		return user, nil
	}
	//???
	return user, model.ErrNotFoundData
}

func (d *dao) deleteUserDB(c context.Context, uid int64) error {
	db := d.db
	result, err := db.Exec(fmt.Sprintf("DELETE FROM user_table WHERE uid = '%v'", uid))
	if err != nil {
		err = fmt.Errorf("exec delete: %w", err)
		return err
	}
	num, err := result.RowsAffected()
	//???
	if num == 0 {
		return model.ErrNotFoundData
	}
	log.Printf("mysql delete user uid = %v, affected = %v\n", uid, num)
	return nil
}

//
func (d *dao) CreateUser(c context.Context, user *model.User) error {
	err := d.createUserDB(c, user)
	if err != nil {
		return err
	}
	return nil
}

//
func (d *dao) UpdateUser(c context.Context, user *model.User) error {
	err := d.updateUserDB(c, user)
	if err != nil {
		return err
	}
	d.delUserCache(c, user.Uid)
	return nil
}
func (d *dao) ReadUser(c context.Context, uid int64) (model.User, error) {
	user := model.User{}
	exist, err := d.existUserCache(c, uid)
	if err != nil {
		return user, nil
	}
	//cache 命中,返回
	if exist {
		if user, err := d.getUserCache(c, uid); err != nil {
			return user, err
		}
		return user, nil
	}
	//cache 没命中,读 DB
	if user, err = d.readUserDB(c, uid); err != nil {
		return user, err
	}
	//回种 cache
	if err = d.setUserCache(c, &user); err != nil {
		return user, err
	}
	//DB 读到的值
	return user, nil
}

func (d *dao) DeleteUser(c context.Context, uid int64) error {
	if err := d.deleteUserDB(c, uid); err != nil {
		return err
	}
	if err := d.delUserCache(c, uid); err != nil {
		return err
	}
	return nil
}
