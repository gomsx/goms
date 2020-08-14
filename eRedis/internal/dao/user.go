package dao

import (
	"context"
	"fmt"
	"log"

	m "github.com/aivuca/goms/eRedis/internal/model"
)

//
func (d *dao) CreateUser(c context.Context, user *m.User) error {
	if err := d.createUserDB(c, user); err != nil {
		err = fmt.Errorf("create user in db: %w", err)
		return err
	}
	return nil
}

// Cache Aside 读策略
func (d *dao) ReadUser(c context.Context, uid int64) (*m.User, error) {
	// 读 cache
	if exist, err := d.existUserCC(c, uid); err != nil {
		// 查询 cache 失败，返回 err
		return nil, err
	} else if exist {
		// 查询 cache 成功，存在条目
		if user, err := d.getUserCC(c, uid); err != nil {
			// 读 cache 失败，返回 err
			err = fmt.Errorf("get user from cc: %w", err)
			return nil, err
		} else {
			// 读 cache 成功，返回 user
			return user, nil
		}
	}
	// 读 DB (cache 没命中)
	if user, err := d.readUserDB(c, uid); err != nil {
		// 读 DB 失败，返回 err
		err = fmt.Errorf("read user from db: %w", err)
		return nil, err
	} else if err = d.setUserCC(c, user); err != nil {
		// 读 DB 成功，回种 cache 失败，返回 err
		log.Printf("faild to set user cc")
		err = fmt.Errorf("set user to cc: %w", err)
		return nil, err
	} else {
		// 读 DB 成功，回种 cache 成功，返回 user
		return user, nil
	}
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
		log.Printf("cache expiration, uid=%v, err=%v", user.Uid, err)
		err = fmt.Errorf("delete user in cc: %w", err)
		return err
	}
	return nil
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
		log.Printf("cache expiration, uid=%v, err=%v", uid, err)
		err = fmt.Errorf("del user in cc: %w", err)
		return err
	}
	return nil
}
