package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/gomodule/redigo/redis"
)

const (
	_createUser = "INSERT INTO user_table VALUES(?,?,?)"
	_updateUser = "UPDATE user_table SET name=?,sex=? WHERE uid=?"
	_readUser   = "SELECT uid,name,sex FROM user_table WHERE uid=?"
	_deleteUser = "DELETE FROM user_table WHERE uid=?"
)

// redis 提供了查询方法　EXISTS，与　GET SET DEL 同级并列，
// 所以先EXISTS再GET的方案.(EXISTS->GET).
// 优于将EXISTS整合再GET中当数据不存在时返回ErrNotFound的方案.(GET{EXISTS,ErrNotFound})

// MySQL 中UPDATE,DELETE 自身可判断是否存在要操作的数(不存在返回ErrNotFound)，
// 所以没必要先通READ判断再操作，而且这样效率也不高．

func (d *dao) existUserCache(c context.Context, uid int64) (bool, error) {
	rd := d.redis
	key := model.GetRedisKey(uid)
	exist, err := redis.Bool(rd.Do("EXISTS", key))
	if err != nil {
		err = fmt.Errorf("cc do EXISTS: %w", err)
		return exist, err
	}
	log.Printf("cc exist=%v key=%v", exist, key)
	return exist, nil
}

func (d *dao) setUserCache(c context.Context, user *model.User) error {
	rd := d.redis
	key := model.GetRedisKey(user.Uid)
	if _, err := rd.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("cc do HMSET: %w", err)
		return err
	}
	log.Printf("cc set key=%v, value=%v", key, user)
	return nil
}

func (d *dao) getUserCache(c context.Context, uid int64) (model.User, error) {
	rd := d.redis
	user := model.User{}
	key := model.GetRedisKey(uid)
	value, err := redis.Values(rd.Do("HGETALL", key))
	if err != nil {
		err = fmt.Errorf("cc do HGETALL: %w", err)
		return user, err
	}
	if err = redis.ScanStruct(value, &user); err != nil {
		err = fmt.Errorf("cc ScanStruct: %w", err)
		return user, err
	}
	log.Printf("cc get key=%v, value=%v", key, user)
	return user, nil
}

func (d *dao) delUserCache(c context.Context, uid int64) error {
	rd := d.redis
	key := model.GetRedisKey(uid)
	if _, err := rd.Do("DEL", key); err != nil {
		err = fmt.Errorf("cc do DEL: %w", err)
		return err
	}
	log.Printf("cc delete key=%v", key)
	return nil
}

func (d *dao) createUserDB(c context.Context, user *model.User) error {
	db := d.db
	result, err := db.Exec(_createUser, user.Uid, user.Name, user.Sex)
	if err != nil {
		err = fmt.Errorf("db exec insert: %w", err)
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	if num == 0 {
		return model.ErrFailedCreateData
	}
	log.Printf("db insert user=%v ", user)
	return nil
}

func (d *dao) updateUserDB(c context.Context, user *model.User) error {
	db := d.db
	result, err := db.Exec(_updateUser, user.Name, user.Sex, user.Uid)
	if err != nil {
		err = fmt.Errorf("db exec update: %w", err)
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	if num == 0 {
		return model.ErrNotFoundData
	}
	log.Printf("db update user=%v, affected=%v ", user, num)
	return nil
}

func (d *dao) readUserDB(c context.Context, uid int64) (model.User, error) {
	db := d.db
	user := model.User{}
	rows, err := db.Query(_readUser, uid)
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("db query: %w", err)
		return user, err
	}
	if rows.Next() {
		if err = rows.Scan(&user.Uid, &user.Name, &user.Sex); err != nil {
			err = fmt.Errorf("db rows scan: %w", err)
			return user, err
		}
		log.Printf("db read user=%v ", user)
		return user, nil
	}
	//???
	return user, model.ErrNotFoundData
}

func (d *dao) deleteUserDB(c context.Context, uid int64) error {
	db := d.db
	result, err := db.Exec(_deleteUser, uid)
	if err != nil {
		err = fmt.Errorf("db exec delete: %w", err)
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	if num == 0 {
		return model.ErrNotFoundData
	}
	log.Printf("db delete user uid=%v, affected=%v ", uid, num)
	return nil
}

//
func (d *dao) CreateUser(c context.Context, user *model.User) error {
	if err := d.createUserDB(c, user); err != nil {
		err = fmt.Errorf("create user in db: %w", err)
		return err
	}
	return nil
}

//
func (d *dao) UpdateUser(c context.Context, user *model.User) error {
	if err := d.updateUserDB(c, user); err != nil {
		err = fmt.Errorf("update user in db: %w", err)
		return err
	}
	if err := d.delUserCache(c, user.Uid); err != nil {
		err = fmt.Errorf("delete user in cc: %w", err)
		return err
	}
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
			err = fmt.Errorf("get user from cc: %w", err)
			return user, err
		}
		return user, nil
	}
	//cache 没命中,读 DB
	if user, err = d.readUserDB(c, uid); err != nil {
		err = fmt.Errorf("read user from db: %w", err)
		return user, err
	}
	//回种 cache
	if err = d.setUserCache(c, &user); err != nil {
		err = fmt.Errorf("set user to cc: %w", err)
		return user, err
	}
	//DB 读到的值
	return user, nil
}

func (d *dao) DeleteUser(c context.Context, uid int64) error {
	if err := d.deleteUserDB(c, uid); err != nil {
		err = fmt.Errorf("del user in db: %w", err)
		return err
	}
	if err := d.delUserCache(c, uid); err != nil {
		err = fmt.Errorf("del user in cc: %w", err)
		return err
	}
	return nil
}
