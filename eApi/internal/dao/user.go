package dao

import (
	"context"
	"fmt"

	m "github.com/aivuca/goms/eApi/internal/model"
	rqid "github.com/aivuca/goms/eApi/internal/pkg/requestid"
)

//
func (d *dao) CreateUser(c context.Context, user *m.User) error {
	if err := d.createUserDB(c, user); err != nil {
		err = fmt.Errorf("create user in db: %w", err)
		return err
	}
	return nil
}

// Cache Aside 写策略(更新)
func (d *dao) UpdateUser(c context.Context, user *m.User) error {
	// 先更新 DB
	if err := d.updateUserDB(c, user); err != nil {
		err = fmt.Errorf("update user in db: %w", err)
		return err
	}
	// 再删除 cache
	if err := d.delUserCC(c, user.Uid); err != nil {
		// 缓存过期
		log.Error().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msgf("cache expiration, uid=%v, err=%v", user.Uid, err)
		err = fmt.Errorf("delete user in cc: %w", err)
		return err
	}
	return nil
}

// Cache Aside 读策略
func (d *dao) ReadUser(c context.Context, uid int64) (*m.User, error) {
	exist, err := d.existUserCC(c, uid)
	if err != nil {
		return nil, err
	}
	//cache 命中,返回
	if exist {
		user, err := d.getUserCC(c, uid)
		if err != nil {
			err = fmt.Errorf("get user from cc: %w", err)
			return nil, err
		}
		return user, nil
	}
	//cache 没命中,读 DB
	user, err := d.readUserDB(c, uid)
	if err != nil {
		err = fmt.Errorf("read user from db: %w", err)
		return nil, err
	}
	//回种 cache
	if err = d.setUserCC(c, user); err != nil {
		// 回中失败
		log.Warn().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", user.Uid).
			Msg("faild to set user cc")
		err = fmt.Errorf("set user to cc: %w", err)
		return nil, err
	}
	//DB 读到的值
	return user, nil
}

// Cache Aside 写策略(删除)
func (d *dao) DeleteUser(c context.Context, uid int64) error {
	// 先删除 DB
	if err := d.deleteUserDB(c, uid); err != nil {
		err = fmt.Errorf("del user in db: %w", err)
		return err
	}
	// 再删除 cache
	if err := d.delUserCC(c, uid); err != nil {
		// 缓存过期
		log.Error().
			Int64("request_id", rqid.GetIdMust(c)).
			Int64("user_id", uid).
			Msgf("cache expiration, uid=%v, err=%v", uid, err)
		err = fmt.Errorf("del user in cc: %w", err)
		return err
	}
	return nil
}
