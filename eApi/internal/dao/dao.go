package dao

import (
	"context"
	"database/sql"

	m "github.com/aivuca/goms/eApi/internal/model"
	lg "github.com/aivuca/goms/eApi/internal/pkg/log"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
)

// Dao dao interface.
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
	//count
	ReadPing(c context.Context, t string) (*m.Ping, error)
	UpdatePing(c context.Context, p *m.Ping) error
	//user
	CreateUser(c context.Context, user *m.User) error
	ReadUser(c context.Context, uid int64) (*m.User, error)
	UpdateUser(c context.Context, user *m.User) error
	DeleteUser(c context.Context, uid int64) error
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis redis.Conn
}

// Log.
var log = lg.Lgd

// New new a Dao.
func New(cfgpath string) (Dao, func(), error) {
	return new(cfgpath)
}

// New new a dao.
func new(cfgpath string) (*dao, func(), error) {
	mdb, cleanDB, err := newDB(cfgpath)
	if err != nil {
		return nil, nil, err
	}
	mcc, _, err := newCC(cfgpath)
	if err != nil {
		cleanDB()
		return nil, nil, err
	}
	log.Info().Msg("dao ok")
	mdao := &dao{db: mdb, redis: mcc}
	return mdao, mdao.Close, nil
}

// Close close the resource.
func (d *dao) Close() {
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	if _, err = d.redis.Do("PING"); err != nil {
		return
	}
	return d.db.PingContext(ctx)
}
