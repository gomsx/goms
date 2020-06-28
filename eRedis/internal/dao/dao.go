package dao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	. "github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/fuwensun/goms/pkg/conf"
)

// Dao dao interface
type Dao interface {
	Close()

	Ping(ctx context.Context) (err error)
	//count
	UpdatePingCount(c context.Context, t PingType, v PingCount) error
	ReadPingCount(c context.Context, t PingType) (PingCount, error)
	//user-cc
	ExistUserCC(c context.Context, uid int64) (bool, error)
	SetUserCC(c context.Context, user *User) error
	GetUserCC(c context.Context, uid int64) (User, error)
	DelUserCC(c context.Context, uid int64) error
	//user-db
	CreateUserDB(c context.Context, user *User) error
	ReadUserDB(c context.Context, uid int64) (User, error)
	UpdateUserDB(c context.Context, user *User) error
	DeleteUserDB(c context.Context, uid int64) error
	//user
	CreateUser(c context.Context, user *User) error
	ReadUser(c context.Context, uid int64) (User, error)
	UpdateUser(c context.Context, user *User) error
	DeleteUser(c context.Context, uid int64) error
}

// dao dao.
type dao struct {
	db    *sql.DB
	redis redis.Conn
}

// dbcfg mysql config.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

//
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

func getDBConfig(cfgpath string) (dbcfg, error) {
	var cfg dbcfg
	path := filepath.Join(cfgpath, "mysql.yml")
	if err := conf.GetConf(path, &cfg); err != nil {
		log.Printf("get db config file: %v", err)
	}
	if cfg.DSN != "" {
		log.Printf("get config db DSN: %v", cfg.DSN)
		return cfg, nil
	}
	if dsn := os.Getenv("MYSQL_SVC_DSN"); dsn != "" {
		cfg.DSN = dsn
		log.Printf("get env db DSN: %v", cfg.DSN)
		return cfg, nil
	}
	err := fmt.Errorf("get db DSN: %w", ErrNotFoundData)
	return cfg, err
}
func getCCConfig(cfgpath string) (cccfg, error) {
	var cfg cccfg
	path := filepath.Join(cfgpath, "redis.yml")
	if err := conf.GetConf(path, &cfg); err != nil {
		log.Printf("get cc config file: %v", err)
	}
	if cfg.Addr != "" {
		log.Printf("get config cc Addr: %v", cfg.Addr)
		return cfg, nil
	}
	if addr := os.Getenv("REDIS_SVC_ADDR"); addr != "" {
		cfg.Addr = addr
		log.Printf("get env cc Addr: %v", cfg.Addr)
		return cfg, nil
	}
	err := fmt.Errorf("get cc Addr: %w", ErrNotFoundData)
	return cfg, err
}

// New new a dao.
func New(cfgpath string) (Dao, func(), error) {
	//cc
	cf, err := getCCConfig(cfgpath)
	if err != nil {
		return nil, nil, err //?
	}
	mcc, err := redis.Dial("tcp", cf.Addr,
		redis.DialPassword(cf.Pass),
	)
	if err != nil {
		log.Panicf("dial cc: %v", err)
	}
	res, err := mcc.Do("PING")
	if err != nil {
		log.Panicf("ping cc: %v", err)
	}
	log.Printf("ping cc res=%v", res)
	//db
	df, err := getDBConfig(cfgpath)
	if err != nil {
		return nil, nil, err //?
	}
	mdb, err := sql.Open("mysql", df.DSN)
	if err != nil {
		log.Panicf("open db: %v", err)
	}
	if err := mdb.Ping(); err != nil {
		log.Panicf("ping db: %v", err)
	}
	log.Printf("ping db err=%v", err)
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
