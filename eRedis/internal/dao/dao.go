package dao

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gomodule/redigo/redis"
)

// DBConfig mysql config.
type DBConfig struct {
	DSN string `yaml:"dsn"`
}

//
type RDConfig struct {
	Name string
	Addr string `yaml:"addr"`
}

var (
	DBconffile = "mysql.yml"
	DSN        = "user:password@/dbname"

	RDconffile = "redis.yml"
	ADDR       = "127.0.0.1:6379"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	//call
	UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error
	ReadPingCount(c context.Context, t model.PingType) (model.PingCount, error)
	//user
	UpdateUserName(c context.Context, uid int64, name string) error
	ReadUserName(c context.Context, uid int64) (name string, err error)

	SetUserCache(c context.Context, user *model.User) error
	GetUserCache(c context.Context, uid int64) (model.User, error)
	DelUserCache(c context.Context, uid int64) error

	CreateUserDB(c context.Context, user *model.User) error
	UpdateUserDB(c context.Context, user *model.User) error
	ReadUserDB(c context.Context, uid int64) (user model.User, err error)

	CreateUser(c context.Context, user *model.User) error
	UpdateUser(c context.Context, user *model.User) error
	ReadUser(c context.Context, uid int64) (model.User, error)
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis redis.Conn
}

// New new a dao.
func New(confpath string) Dao {

	//db
	var dc DBConfig
	pathname := filepath.Join(confpath, DBconffile)
	if err := conf.GetConf(pathname, &dc); err != nil {
		log.Printf("failed to get db config file! error: %v", err)
	}

	if dc.DSN != "" {
		DSN = dc.DSN
	}
	log.Printf("db DSN: %v", DSN)

	mdb, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Panicf("failed to open db! error: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("failed to ping db! error: %v", err)
	}

	//rd
	var rc RDConfig
	pathname = filepath.Join(confpath, RDconffile)
	if err := conf.GetConf(pathname, &rc); err != nil {
		log.Printf("failed to get rc config file! error: %v", err)
	}

	if rc.Addr != "" {
		ADDR = rc.Addr
	}
	log.Printf("rc addr: %v", ADDR)

	mrd, err := redis.Dial("tcp", ADDR)
	if err != nil {
		log.Panicf("failed to conn rc! error: %v", err)
	}
	if _, err := mrd.Do("PING"); err != nil {
		log.Panicf("failed to ping rc! error: %v", err)
	}

	return &dao{
		db:    mdb,
		redis: mrd,
	}
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
