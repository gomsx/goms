package dao

import (
	"context"
	"fmt"

	. "github.com/fuwensun/goms/eLog/internal/model"

	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
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

func (d *dao) ExistUserCC(c context.Context, uid int64) (bool, error) {
	cc := d.redis
	key := GetRedisKey(uid)
	exist, err := redis.Bool(cc.Do("EXISTS", key))
	if err != nil {
		err = fmt.Errorf("cc do EXISTS: %w", err)
		return exist, err
	}
	log.Debug().Str("key", key).Msgf("cc %v exist", exist)
	return exist, nil
}

func (d *dao) SetUserCC(c context.Context, user *User) error {
	cc := d.redis
	key := GetRedisKey(user.Uid)
	if _, err := cc.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("cc do HMSET: %w", err)
		return err
	}
	log.Debug().Str("key", key).Msg("cc set user")
	return nil
}

func (d *dao) GetUserCC(c context.Context, uid int64) (*User, error) {
	cc := d.redis
	user := &User{}
	key := GetRedisKey(uid)
	value, err := redis.Values(cc.Do("HGETALL", key))
	if err != nil {
		err = fmt.Errorf("cc do HGETALL: %w", err)
		return user, err
	}
	if err = redis.ScanStruct(value, user); err != nil {
		err = fmt.Errorf("cc ScanStruct: %w", err)
		return user, err
	}
	log.Debug().Str("key", key).Msg("cc get user")
	return user, nil
}

func (d *dao) DelUserCC(c context.Context, uid int64) error {
	cc := d.redis
	key := GetRedisKey(uid)
	if _, err := cc.Do("DEL", key); err != nil {
		err = fmt.Errorf("cc do DEL: %w", err)
		return err
	}
	log.Debug().Str("key", key).Msg("cc delete user")
	return nil
}

func (d *dao) CreateUserDB(c context.Context, user *User) error {
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

	log.Info().Int64("uid", user.Uid).Msg("db insert user")
	log.Debug().Int64("rows", num).Msg("db insert user")
	return nil
}

func (d *dao) ReadUserDB(c context.Context, uid int64) (*User, error) {
	db := d.db
	user := &User{}
	rows, err := db.Query(_readUser, uid)
	if err != nil {
		err = fmt.Errorf("db query: %w", err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&user.Uid, &user.Name, &user.Sex); err != nil {
			err = fmt.Errorf("db rows scan: %w", err)
			return nil, err
		}
		if rows.Next() {
			// uid 重复
			log.Error().Int64("uid", uid).Msg("db read multiple uid")
		}
		log.Debug().Int64("uid", uid).Msg("db read user")
		return user, nil
	}
	//not found
	log.Debug().Int64("uid", uid).Msg("db not found user")
	return user, nil
}

func (d *dao) UpdateUserDB(c context.Context, user *User) error {
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
	log.Info().Int64("uid", user.Uid).Msg("db update user")
	log.Debug().Int64("rows", num).Msg("db update user")
	return nil
}

func (d *dao) DeleteUserDB(c context.Context, uid int64) error {
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
	log.Info().Int64("uid", uid).Msg("db delete user")
	log.Debug().Int64("rows", num).Msg("db delete user")
	return nil
}

//
func (d *dao) CreateUser(c context.Context, user *User) error {
	if err := d.CreateUserDB(c, user); err != nil {
		err = fmt.Errorf("create user in db: %w", err)
		return err
	}
	return nil
}

// Cache Aside 写策略(更新)
func (d *dao) UpdateUser(c context.Context, user *User) error {
	// 先更新 DB
	if err := d.UpdateUserDB(c, user); err != nil {
		err = fmt.Errorf("update user in db: %w", err)
		return err
	}
	// 再删除 cache
	if err := d.DelUserCC(c, user.Uid); err != nil {
		// 缓存过期
		log.Error().Msgf("cache expiration, uid=%v, err=%v", user.Uid, err)
		err = fmt.Errorf("delete user in cc: %w", err)
		return err
	}
	return nil
}

// Cache Aside 读策略
func (d *dao) ReadUser(c context.Context, uid int64) (*User, error) {
	exist, err := d.ExistUserCC(c, uid)
	if err != nil {
		return nil, err
	}
	//cache 命中,返回
	if exist {
		user, err := d.GetUserCC(c, uid)
		if err != nil {
			err = fmt.Errorf("get user from cc: %w", err)
			return nil, err
		}
		return user, nil
	}
	//cache 没命中,读 DB
	user, err := d.ReadUserDB(c, uid)
	if err != nil {
		err = fmt.Errorf("read user from db: %w", err)
		return nil, err
	}
	//回种 cache
	if err = d.SetUserCC(c, user); err != nil {
		// 回中失败
		log.Warn().Int64("uid", user.Uid).Msg("faild to set user cc")
		err = fmt.Errorf("set user to cc: %w", err)
		return nil, err
	}
	//DB 读到的值
	return user, nil
}

// Cache Aside 写策略(删除)
func (d *dao) DeleteUser(c context.Context, uid int64) error {
	// 先删除 DB
	if err := d.DeleteUserDB(c, uid); err != nil {
		err = fmt.Errorf("del user in db: %w", err)
		return err
	}
	// 再删除 cache
	if err := d.DelUserCC(c, uid); err != nil {
		// 缓存过期
		log.Error().Msgf("cache expiration, uid=%v, err=%v", uid, err)
		err = fmt.Errorf("del user in cc: %w", err)
		return err
	}
	return nil
}
