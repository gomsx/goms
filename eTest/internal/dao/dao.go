package dao

import (
	"context"
	"database/sql"

	m "github.com/aivuca/goms/eTest/internal/model"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

// Dao dao interface.
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	//count
	ReadPing(ctx context.Context, t string) (*m.Ping, error)
	UpdatePing(ctx context.Context, p *m.Ping) error
	//user
	CreateUser(ctx context.Context, user *m.User) error
	ReadUser(ctx context.Context, uid int64) (*m.User, error)
	UpdateUser(ctx context.Context, user *m.User) error
	DeleteUser(ctx context.Context, uid int64) error
}

// dao dao struct.
type dao struct {
	db    *sql.DB
	redis redis.Conn
}

// New new Dao and return.
func New(cfgpath string) (Dao, func(), error) {
	return new(cfgpath)
}

// New new dao and return.
func new(cfgpath string) (*dao, func(), error) {
	mdb, cleanDB, err := newDB(cfgpath)
	if err != nil {
		return nil, nil, err
	}
	log.Infof("db ok")
	mcc, _, err := newCC(cfgpath)
	if err != nil {
		cleanDB()
		return nil, nil, err
	}
	log.Infof("cc ok")
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
