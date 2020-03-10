package dao

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gomodule/redigo/redis"
)

// DBConfig mysql config.
type DBConfig struct {
	DSN string `yaml:"dsn"`
}

//
type CCConfig struct {
	Name string
	Addr string `yaml:"addr"`
}

var (
	DBconfigfile = "mysql.yml"
	DBDSN        = "user:password@/dbname"

	CCconfigfile = "redis.yml"
	CCADDR       = "127.0.0.1:6379"
)

// Dao dao interface
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
	//call
	UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error
	ReadPingCount(c context.Context, t model.PingType) (model.PingCount, error)

	//user-cc
	existUserCache(c context.Context, uid int64) (bool, error)
	setUserCache(c context.Context, user *model.User) error
	getUserCache(c context.Context, uid int64) (model.User, error)
	delUserCache(c context.Context, uid int64) error
	//user-db
	createUserDB(c context.Context, user *model.User) error
	updateUserDB(c context.Context, user *model.User) error
	readUserDB(c context.Context, uid int64) (user model.User, err error)
	deleteUserDB(c context.Context, uid int64) error
	//user
	CreateUser(c context.Context, user *model.User) error
	UpdateUser(c context.Context, user *model.User) error
	ReadUser(c context.Context, uid int64) (model.User, error)
	DeleteUser(c context.Context, uid int64) error
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
	pathname := filepath.Join(confpath, DBconfigfile)
	if err := conf.GetConf(pathname, &dc); err != nil {
		log.Printf("get db config file: %v", err)
	}
	if dc.DSN != "" {
		DBDSN = dc.DSN
	}
	log.Printf("db DSN: %v", DBDSN)

	mdb, err := sql.Open("mysql", DBDSN)
	if err != nil {
		log.Panicf("open db: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("ping db: %v", err)
	}
	//cc
	var cc CCConfig
	pathname = filepath.Join(confpath, CCconfigfile)
	if err := conf.GetConf(pathname, &cc); err != nil {
		log.Printf("get cc config file: %v", err)
	}
	if cc.Addr != "" {
		CCADDR = cc.Addr
	}
	log.Printf("cc addr: %v", CCADDR)

	mrd, err := redis.Dial("tcp", CCADDR)
	if err != nil {
		log.Panicf("dial redis: %v", err)
	}
	if _, err := mrd.Do("PING"); err != nil {
		log.Panicf("ping redis: %v", err)
	}
	if _, err := mrd.Do("FLUSHDB"); err != nil {
		log.Panicf("flush redis: %v", err)
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
