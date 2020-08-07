package dao

import (
	"context"
	"fmt"

	m "github.com/aivuca/goms/eApi/internal/model"
	"github.com/aivuca/goms/eApi/internal/pkg/reqid"

	"github.com/gomodule/redigo/redis"
)

func (d *dao) existUserCC(c context.Context, uid int64) (bool, error) {
	cc := d.redis
	key := m.GetRedisKey(uid)
	exist, err := redis.Bool(cc.Do("EXISTS", key))
	if err != nil {
		err = fmt.Errorf("cc do EXISTS: %w", err)
		return exist, err
	}
	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", uid).
		Str("key", key).
		Msgf("cc %v exist", exist)
	return exist, nil
}

func (d *dao) setUserCC(c context.Context, user *m.User) error {
	cc := d.redis
	key := m.GetRedisKey(user.Uid)
	if _, err := cc.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("cc do HMSET: %w", err)
		return err
	}
	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", user.Uid).
		Str("key", key).
		Msg("cc set user")
	return nil
}

func (d *dao) getUserCC(c context.Context, uid int64) (*m.User, error) {
	cc := d.redis
	user := &m.User{}
	key := m.GetRedisKey(uid)
	value, err := redis.Values(cc.Do("HGETALL", key))
	if err != nil {
		err = fmt.Errorf("cc do HGETALL: %w", err)
		return user, err
	}
	if err = redis.ScanStruct(value, user); err != nil {
		err = fmt.Errorf("cc ScanStruct: %w", err)
		return user, err
	}
	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", uid).
		Str("key", key).
		Msg("cc get user")
	return user, nil
}

func (d *dao) delUserCC(c context.Context, uid int64) error {
	cc := d.redis
	key := m.GetRedisKey(uid)
	if _, err := cc.Do("DEL", key); err != nil {
		err = fmt.Errorf("cc do DEL: %w", err)
		return err
	}
	log.Debug().
		Int64("request_id", reqid.GetIdMust(c)).
		Int64("user_id", uid).
		Str("key", key).
		Msg("cc delete user")
	return nil
}
