package dao

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gomodule/redigo/redis"
)

// dbcfg mysql config.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

//
type cccfg struct {
	Name string
	Addr string `yaml:"addr"`
}

var (
	dbcfgfile = "mysql.yml"
	dbDSN     = "user:password@/dbname"

	cccfgfile = "redis.yml"
	ccADDR    = "127.0.0.1:6379"
)

// Dao dao interface
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
	//call
	UpdatePingCount(c context.Context, t model.PingType, v model.PingCount) error
	ReadPingCount(c context.Context, t model.PingType) (model.PingCount, error)
	//user-cc
	existUserCC(c context.Context, uid int64) (bool, error)
	setUserCC(c context.Context, user *model.User) error
	getUserCC(c context.Context, uid int64) (model.User, error)
	delUserCC(c context.Context, uid int64) error
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
func New(cfgpath string) (Dao, func(), error) {
	//db
	var df dbcfg
	path := filepath.Join(cfgpath, dbcfgfile)
	if err := conf.GetConf(path, &df); err != nil {
		log.Printf("get db config file: %v", err)
		// err = fmt.Errorf("get db config file: %w", err)
		// return nil, nil, err
	}
	if df.DSN != "" {
		dbDSN = df.DSN
		log.Printf("get config db DSN: %v", df.DSN)
	}
	if dsn := os.Getenv("MYSQL_SVC_DSN"); dsn != "" {
		dbDSN = dsn
		log.Printf("get env db DSN: %v", dsn)
	}
	mdb, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Panicf("open db: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("ping db: %v", err)
	}
	//cc
	var cf cccfg
	path = filepath.Join(cfgpath, cccfgfile)
	if err := conf.GetConf(path, &cf); err != nil {
		log.Printf("get cc config file: %v", err)
		// err = fmt.Errorf("get cc config file: %w", err)
		// return nil, nil, err
	}
	if cf.Addr != "" {
		ccADDR = cf.Addr
		log.Printf("get config cc ADDR: %v", cf.Addr)
	}
	if addr := os.Getenv("REDIS_SVC_ADDR"); addr != "" {
		ccADDR = addr
		log.Printf("get env cc ADDR: %v", addr)
	}
	mcc, err := redis.Dial("tcp", ccADDR)
	if err != nil {
		log.Panicf("dial cc: %v", err)
	}
	if _, err := mcc.Do("PING"); err != nil {
		log.Panicf("ping cc: %v", err)
	}
	//
	d := &dao{
		db:    mdb,
		redis: mcc,
	}
	return d, d.Close, nil
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
